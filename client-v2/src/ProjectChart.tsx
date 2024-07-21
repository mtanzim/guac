import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";

import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Card, CardContent, CardHeader, CardTitle } from "./components/ui/card";
import { StatsData } from "./data-types";

export function ProjectChart({
  projectDurations,
}: {
  projectDurations: StatsData["projectStats"]["durations"];
}) {
  const chartConfig = {
    hours: {
      label: "Hours",
    },
  } satisfies ChartConfig;

  const chartData = projectDurations.map((d) => ({
    hours: (d.minutes / 60).toFixed(2),
    project: d.project,
  }));

  return (
    <Card>
      <CardHeader>
        <CardTitle>Projects worked on</CardTitle>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="min-h-[200px] w-full">
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <YAxis dataKey={"hours"} />
            <XAxis
              dataKey="project"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
            />
            <ChartTooltip content={<ChartTooltipContent />} />
            <Bar dataKey={"hours"} radius={4} />;
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
