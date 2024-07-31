import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";

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
import { toCustomDateStr, toCustomDateStrWithYear } from "./utils";
import { SkeletonChart } from "./SkeletonChart";

export function DailyChart({
  dailyDuration,
  loading,
}: {
  dailyDuration?: StatsData["dailyDuration"];
  loading: boolean;
}) {
  const chartConfig = {
    hours: {
      label: "Hours",
    },
  } satisfies ChartConfig;

  const dateFormatter =
    (dailyDuration || [])?.length > Math.floor(365 / 2)
      ? toCustomDateStrWithYear
      : toCustomDateStr;

  const chartData = (dailyDuration || []).map((d) => ({
    hours: (d.minutes / 60).toFixed(2),
    date: dateFormatter(d.date),
  }));
  if (loading) {
    return <SkeletonChart />;
  }
  return (
    <Card>
      <CardHeader>
        <CardTitle>Time coding</CardTitle>
        <CardDescription>Hours spent</CardDescription>
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
              tickFormatter={dateFormatter}
            />
            <ChartTooltip content={<ChartTooltipContent />} />
            <Bar dataKey={"hours"} radius={4} />;
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
