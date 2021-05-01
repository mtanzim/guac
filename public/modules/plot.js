export function plotLangPie(percentages, divName) {
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
    height: 800,
    width: 800,
  };
  Plotly.newPlot(divName, data, layout);
}
