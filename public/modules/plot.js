export function plotLangPie(divName, { percentages, startDate, endDate }) {
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
      // marker: {
      //   colors: percentages.map((d) => d.color),
      // },
    },
  ];

  const layout = {
    title: `Languages Used: ${startDate} to ${endDate}`,
    font: { size: 12 },
    height: 600,
    width: 600,
  };
  const config = { responsive: true };
  Plotly.newPlot(divName, data, layout, config);
}

export function plotDailyDur(divName, { dailyDuration, startDate, endDate }) {
  var data = [
    {
      x: dailyDuration.map((d) => d.date),
      y: dailyDuration.map((d) => (d.minutes / 60).toFixed(0)),
      type: "bar",
    },
  ];
  const layout = {
    title: `Daily coding activity: ${startDate} to ${endDate}`,
    xaxis: { title: "Date" },
    yaxis: { title: "Hours spent" },
    font: { size: 12 },
    height: 600,
    width: 600,
  };
  const config = { responsive: true };
  Plotly.newPlot(divName, data, layout, config);
}

export function plotLangDur(divName, { langDur, startDate, endDate }) {
  var data = [
    {
      x: langDur.map((d) => d.language),
      y: langDur.map((d) => (d.minutes / 60).toFixed(0)),
      type: "bar",
    },
  ];
  const layout = {
    title: `Time spent on languages: ${startDate} to ${endDate}`,
    xaxis: { title: "Language" },
    yaxis: { title: "Hours spent" },
    font: { size: 12 },
    height: 600,
    width: 600,
  };
  const config = { responsive: true };
  Plotly.newPlot(divName, data, layout, config);
}
