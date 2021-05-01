import {
  plotLangPie,
  plotDailyDur,
  plotLangDur,
  plotProjDur,
} from "./modules/plot.js";

const TOKEN_KEY = "WakaToken";
// TODO: logout
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
      return token;
    });
}

function plotData(data) {
  console.log(data);
  const {
    startDate,
    endDate,
    dailyDuration,
    projectStats,
    languageStats,
  } = data;
  const { percentages, durations: langDur } = languageStats;
  const { durations: projDur } = projectStats;
  plotLangPie("lang-pie", { percentages, startDate, endDate });
  plotLangDur("lang-dur", { langDur, startDate, endDate });
  plotDailyDur("daily-dur", { dailyDuration, startDate, endDate });
  plotProjDur("proj-dur", { projDur, startDate, endDate });
}

function fetchData(token) {
  return fetch("/api/v1/data", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }).then((res) => {
    if (res.status === 200) {
      const loginForm = document.getElementById("login-form");
      loginForm.remove();
      return res.json();
    }
    throw new Error("Failed to get data");
  });
}

function initWaka() {
  const curToken = window.localStorage.getItem(TOKEN_KEY);
  if (curToken) {
    return fetchData(curToken)
      .then(plotData)
      .catch((err) => {
        document.getElementById("error").innerText = err.message;
      });
  }

  const loginBtn = document.getElementById("login-btn");
  loginBtn.onclick = () =>
    login()
      .then((token) => fetchData(token))
      .then(plotData)
      .catch((err) => {
        document.getElementById("error").innerText = err.message;
      });
}

window.initWaka = initWaka;
