import { Bar, BarChart, CartesianGrid, Cell, XAxis, YAxis } from "recharts";

import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Card, CardContent, CardHeader, CardTitle } from "./components/ui/card";
import { StatsData } from "./data-types";
import { getColor } from "./utils";

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

  const chartData = languageDurations.map((d) => ({
    hours: (d.minutes / 60).toFixed(2),
    language: d.language,
  }));

  return (
    <Card>
      <CardHeader>
        <CardTitle>Languages Used</CardTitle>
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
