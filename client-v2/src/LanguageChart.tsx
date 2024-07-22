import { Bar, BarChart, CartesianGrid, Cell, XAxis, YAxis } from "recharts";

import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./components/ui/card";
import { StatsData } from "./data-types";
import { getColor, TOP_N_LANGUAGES } from "./utils";

export function LanguageChart({
  languageDurations,
}: {
  languageDurations: StatsData["languageStats"]["durations"];
}) {
  const chartConfig = {
    hours: {
      label: "Hours",
    },
  } satisfies ChartConfig;

  const chartData = languageDurations
    .slice()
    .sort((a, b) => b.minutes - a.minutes)
    .slice(0, TOP_N_LANGUAGES)
    .map((d) => ({
      hours: (d.minutes / 60).toFixed(2),
      language: d.language,
    }));

  const restMinutes = languageDurations
    .slice()
    .sort((a, b) => b.minutes - a.minutes)
    .slice(TOP_N_LANGUAGES)
    .reduce((acc, cur) => acc + cur.minutes, 0);
  chartData.push({ hours: (restMinutes / 60).toFixed(2), language: "Rest" });

  return (
    <Card>
      <CardHeader>
        <CardTitle>Languages Used</CardTitle>
        <CardDescription>Hours spent</CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="min-h-[200px] w-full">
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <YAxis dataKey={"hours"} />
            <XAxis
              dataKey="language"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              // tickFormatter={toCustomDateStr}
            />
            <ChartTooltip content={<ChartTooltipContent />} />
            <Bar dataKey={"hours"} radius={4}>
              {chartData.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={getColor(entry.language)} />
              ))}
            </Bar>
            ;
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
