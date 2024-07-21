import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";

import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Card, CardContent, CardHeader, CardTitle } from "./components/ui/card";
import { StatsData } from "./data-types";
import { toCustomDateStr } from "./utils";

export function DailyChart({
  dailyDuration,
}: {
  dailyDuration: StatsData["dailyDuration"];
}) {
  console.log({ dailyDuration });
  const chartConfig = {
    hours: {
      label: "Hours",
    },
  } satisfies ChartConfig;

  const chartData = dailyDuration.map((d) => ({
    hours: (d.minutes / 60).toFixed(2),
    date: d.date,
  }));

  return (
    <Card>
      <CardHeader>
        <CardTitle>Daily time spent coding</CardTitle>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="min-h-[200px] w-full">
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <YAxis dataKey={"hours"} />
            <XAxis
              dataKey="date"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={toCustomDateStr}
            />
            <ChartTooltip content={<ChartTooltipContent />} />
            <Bar dataKey={"hours"} radius={4} />;
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
