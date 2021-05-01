import { plotLangPie } from "./modules/plot.js";

const TOKEN_KEY = "WakaToken";

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
  const { percentages } = languageStats;
  plotLangPie(percentages, "lang-pie");
}

function fetchData(token) {
  const loginForm = document.getElementById("login-form");
  loginForm.remove();
  return fetch("/api/v1/data", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }).then((res) => res.json());
}

function initWaka() {
  const curToken = window.localStorage.getItem(TOKEN_KEY);
  if (curToken) {
    return fetchData(curToken).then(plotData);
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
