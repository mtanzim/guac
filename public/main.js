import { plotLangPie } from "./modules/plot.js";

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

window.initWaka = () =>
  login()
    .then((token) => {
      const loginForm = document.getElementById("login-form");
      loginForm.remove();
      return fetch("/api/v1/data", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
    })
    .then((res) => res.json())
    .then(plotData)
    .catch((err) => {
      document.getElementById("error").innerText = err.message;
    });
