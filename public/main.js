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
  return Math.round(Math.abs((new Date(start) - new Date(end)) / oneDay)) + 1;
}

async function plotData(data) {
  const { startDate, endDate, dailyDuration, projectStats, languageStats } =
    data;
  const subtitle = document.getElementById("subtitle");

  const start = formatDate(startDate);
  const end = formatDate(endDate);
  const diff = daysBetween(start, end);

  subtitle.innerHTML = `
  <small>${start} to ${end}</small>
  <br/>
  <small>${diff} days</small>
  `;

  const { percentages, durations: langDur } = languageStats;
  const { durations: projDur } = projectStats;
  plotLangDur("lang-dur", { langDur });
  plotDailyDur("daily-dur", { dailyDuration });
  plotProjDur("proj-dur", { projDur });
  const colors = await getColors();
  plotLangPie("lang-pie", { percentages, colors });
}

function fetchData(token, start, end) {
  const url =
    start && end ? `/api/v1/data?start=${start}&end=${end}` : "/api/v1/data";
  return fetch(url, {
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
  const controlBtns = document.getElementById("control-btns");
  controlBtns.style.display = "block";
  const plots = document.getElementById("plots");
  plots.style.display = "grid";
}

function getDateRange(days) {
  const formatDateForReq = (date) => {
    return date.toISOString().split("T")[0];
  };

  const endDate = new Date();
  endDate.setDate(endDate.getDate() - 1);
  const ending = formatDateForReq(endDate);
  const startDate = new Date();
  startDate.setDate(startDate.getDate() - days);
  const starting = formatDateForReq(startDate);
  return { starting, ending };
}

function attachTimeRangeControl(token) {
  const oneWkButton = document.getElementById("one-wk");
  const twoWkButton = document.getElementById("two-wk");
  const oneMonthButton = document.getElementById("last-m");
  const threeMonthButton = document.getElementById("last-3m");
  const allTimeButton = document.getElementById("all-time");

  const generateHandler = (days, curBtn) => {
    const { starting, ending } = getDateRange(days);
    return () => {
      const allCtrlBtns = document.querySelectorAll(".control");
      allCtrlBtns.forEach((btn) => {
        if (curBtn.id === btn.id) {
          btn.classList.add("control-active");
        } else {
          btn.classList.remove("control-active");
        }
      });
      showPlots(token, starting, ending);
    };
  };

  oneWkButton.onclick = generateHandler(7, oneWkButton);
  twoWkButton.onclick = generateHandler(14, twoWkButton);
  oneMonthButton.onclick = generateHandler(30, oneMonthButton);
  threeMonthButton.onclick = generateHandler(30 * 3, threeMonthButton);
  allTimeButton.onclick = generateHandler(365 * 3, allTimeButton);
}

function showPlots(token, start, end) {
  return fetchData(token, start, end)
    .then(plotData)
    .catch((err) => {
      document.getElementById("error").innerText = err.message;
    });
}

const DEFAULT_DAY_RANGE = 7;
const { starting: DEFAULT_START, ending: DEFAULT_END } =
  getDateRange(DEFAULT_DAY_RANGE);

function showLoginForm() {
  const plots = document.getElementById("plots");
  plots.style.display = "none";
  const loginForm = document.getElementById("login-form");
  loginForm.style.display = "block";
  document.getElementById("error").innerText = "";
  const controlBtns = document.getElementById("control-btns");
  controlBtns.style.display = "none";
  const sub = document.getElementById("subtitle");
  sub.innerHTML = "";
  const loginBtn = document.getElementById("login-btn");
  loginBtn.onclick = () => {
    login().then((token) => {
      showPlots(token, DEFAULT_START, DEFAULT_END);
      attachTimeRangeControl(token);
    });
  };
}

function initWaka() {
  const curToken = window.localStorage.getItem(TOKEN_KEY);
  if (curToken) {
    hideLoginForm();
    showPlots(curToken, DEFAULT_START, DEFAULT_END);
    attachTimeRangeControl(curToken);
  } else {
    showLoginForm();
  }
}

// TODO: this is a mess; never use vanillaJS again
// TODO: allow selecting start and end days
window.initWaka = initWaka;
window.logoutWaka = logout;
