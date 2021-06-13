import {
  plotLangPie,
  plotDailyDur,
  plotLangDur,
  plotProjDur,
} from "./modules/plot.js";

async function getColors() {
  const res = await fetch("colors.json");
  return res.json();
}

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

      window.localStorage.setItem(TOKEN_KEY, token);
      hideLoginForm();
      return token;
    });
}

function formatDate(date) {
  var options = {
    year: "numeric",
    month: "short",
    day: "numeric",
    timeZone: "UTC",
  };
  return new Date(date).toLocaleDateString("en-US", options);
}

function daysBetween(start, end) {
  const oneDay = 24 * 60 * 60 * 1000; // hours*minutes*seconds*milliseconds
  return Math.round(Math.abs((new Date(start) - new Date(end)) / oneDay));
}

async function plotData(data) {
  const { startDate, endDate, dailyDuration, projectStats, languageStats } =
    data;
  const subtitle = document.getElementById("subtitle");

  const start = formatDate(startDate);
  const end = formatDate(endDate);
  const diff = daysBetween(start, end);

  subtitle.innerHTML = `
  <code>${start} to ${end}</code>
  <br/>
  <code>${diff} days</code>
  `;

  const { percentages, durations: langDur } = languageStats;
  const { durations: projDur } = projectStats;
  const colors = await getColors();
  plotLangPie("lang-pie", { percentages, colors });
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
  logoutBtn.style.display = "block";
  const plots = document.getElementById("plots");
  plots.style.display = "grid";
}

function showLoginForm() {
  const loginForm = document.getElementById("login-form");
  loginForm.style.display = "block";
  document.getElementById("error").innerText = "";
  const logoutBtn = document.getElementById("logout-btn");
  logoutBtn.style.display = "none";
  const sub = document.getElementById("subtitle");
  sub.innerHTML = "";
  const loginBtn = document.getElementById("login-btn");
  loginBtn.onclick = () =>
    login()
      .then((token) => fetchData(token))
      .then(plotData)
      .catch((err) => {
        document.getElementById("error").innerText = err.message;
      });
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
}

// TODO: this is a mess; never use vanillaJS again
// TODO: allow selecting start and end days
window.initWaka = initWaka;
window.logoutWaka = logout;
