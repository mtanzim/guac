import {
  plotLangPie,
  plotDailyDur,
  plotLangDur,
  plotProjDur,
} from "./modules/plot.js";

const TOKEN_KEY = "WakaToken";

function logout() {
  showLoginForm();
  window.localStorage.clear(TOKEN_KEY);
  const plots = document.getElementById("plots");
  plots.style.display = "none";
}

function login() {
  const username = document.getElementById("username").value;
  const password = document.getElementById("pass").value;
  return fetch("/api/v1/login", {
    method: "POST",
    body: JSON.stringify({
      username,
      password,
    }),
  })
    .then((res) => res.json())
    .then((data) => {
      const { token } = data;
      if (!token) {
        throw new Error("Unable to login");
      }
      const plots = document.getElementById("plots");
      plots.style.display = "grid";
      window.localStorage.setItem(TOKEN_KEY, token);
      hideLoginForm();
      return token;
    });
}

function formatDate(date) {
  var options = { year: "numeric", month: "short", day: "numeric" };
  return new Date(date).toLocaleDateString("en-US", options);
}

function daysBetween(start, end) {
  const oneDay = 24 * 60 * 60 * 1000; // hours*minutes*seconds*milliseconds
  return Math.round(Math.abs((new Date(start) - new Date(end)) / oneDay));
}

function plotData(data) {
  const {
    startDate,
    endDate,
    dailyDuration,
    projectStats,
    languageStats,
  } = data;
  const subtitle = document.getElementById("subtitle");

  const start = formatDate(startDate);
  const end = formatDate(endDate);
  const diff = daysBetween(start, end);

  subtitle.innerHTML = `
  <h3>${start} to ${end}</h3>
  <h3>${diff} days</h3>
  `;

  const { percentages, durations: langDur } = languageStats;
  const { durations: projDur } = projectStats;
  plotLangPie("lang-pie", { percentages });
  plotLangDur("lang-dur", { langDur });
  plotDailyDur("daily-dur", { dailyDuration });
  plotProjDur("proj-dur", { projDur });
}

function fetchData(token) {
  return fetch("/api/v1/data", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }).then((res) => {
    if (res.status === 200) {
      return res.json();
    }
    logout();
    throw new Error("Failed to get data");
  });
}

function hideLoginForm() {
  document.getElementById("username").value = "";
  document.getElementById("pass").value = "";
  const loginForm = document.getElementById("login-form");
  loginForm.style.display = "none";
  document.getElementById("error").innerText = "";
  const logoutBtn = document.getElementById("logout-btn");
  logoutBtn.style.visibility = "visible";
}

function showLoginForm() {
  const loginForm = document.getElementById("login-form");
  loginForm.style.display = "block";
  document.getElementById("error").innerText = "";
  const logoutBtn = document.getElementById("logout-btn");
  logoutBtn.style.visibility = "hidden";
  const sub = document.getElementById("subtitle");
  sub.innerHTML = "";
}

function initWaka() {
  const curToken = window.localStorage.getItem(TOKEN_KEY);
  if (curToken) {
    hideLoginForm();
    return fetchData(curToken)
      .then(plotData)
      .catch((err) => {
        document.getElementById("error").innerText = err.message;
      });
  }

  showLoginForm();
  const loginBtn = document.getElementById("login-btn");
  loginBtn.onclick = () =>
    login()
      .then((token) => fetchData(token))
      .then(plotData)
      .catch((err) => {
        document.getElementById("error").innerText = err.message;
      });
}

// TODO: this is a mess; never use vanillaJS again
// TODO: allow selecting start and end days
window.initWaka = initWaka;
window.logoutWaka = logout;
