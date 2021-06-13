const WIDTH = 800;
const HEIGHT = WIDTH;
const font = {
  family: "Roboto, sans-serif",
  size: 14,
  color: "#7f7f7f",
};
export function plotLangPie(divName, { percentages, colors }) {
  const data = [
    {
      labels: percentages.map((d) => d.language),
      values: percentages.map((d) => d.percentage.toFixed(1)),
      type: "pie",
      hole: 0.7,
      textinfo: "label",
      textposition: "outside",
      automargin: true,
      // TODO: colors
      marker: {
        colors: percentages.map((d) => colors[d.language]?.color),
      },
    },
  ];

  const layout = {
    title: `Languages Used`,
    height: HEIGHT,
    width: WIDTH,
    ...font,
  };
  const config = { responsive: true };
  Plotly.newPlot(divName, data, layout, config);
}

export function plotDailyDur(divName, { dailyDuration }) {
  var data = [
    {
      x: dailyDuration.map((d) => d.date),
      y: dailyDuration.map((d) => (d.minutes / 60).toFixed(1)),
      type: "bar",
    },
  ];
  const layout = {
    title: `Daily time spent coding`,
    xaxis: { title: "Date" },
    yaxis: { title: "Hours spent" },
    height: HEIGHT,
    width: WIDTH,
    ...font,
  };
  const config = { responsive: true };
  setTimeout(() => Plotly.newPlot(divName, data, layout, config), 50);
}

export function plotLangDur(divName, { langDur }) {
  var data = [
    {
      x: langDur.map((d) => d.language),
      y: langDur.map((d) => (d.minutes / 60).toFixed(1)),
      type: "bar",
    },
  ];
  const layout = {
    title: `Time spent on languages`,
    xaxis: { title: "Language" },
    yaxis: { title: "Hours spent" },
    height: HEIGHT,
    width: WIDTH,
    ...font,
  };
  const config = { responsive: true };
  Plotly.newPlot(divName, data, layout, config);
}

export function plotProjDur(divName, { projDur }) {
  var data = [
    {
      x: projDur.map((d) => d.project),
      y: projDur.map((d) => (d.minutes / 60).toFixed(1)),
      type: "bar",
    },
  ];
  const layout = {
    title: `Time spent on projects`,
    xaxis: { title: "Project" },
    yaxis: { title: "Hours spent" },
    height: HEIGHT,
    width: WIDTH,
    ...font,
  };
  const config = { responsive: true };
  Plotly.newPlot(divName, data, layout, config);
}
