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
import { TOP_N_PROJECTS } from "./utils";
import { useMemo } from "react";

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

  const [chartData, maxY] = useMemo(() => {
    const d = projectDurations
      .slice()
      .sort((a, b) => b.minutes - a.minutes)
      .slice(0, TOP_N_PROJECTS)
      .map((d) => ({
        hours: (d.minutes / 60).toFixed(2),
        project: d.project,
      }));

    const restMinutes = projectDurations
      .slice()
      .sort((a, b) => b.minutes - a.minutes)
      .slice(TOP_N_PROJECTS)
      .reduce((acc, cur) => acc + cur.minutes, 0);
    d.push({ hours: (restMinutes / 60).toFixed(2), project: "Rest" });
    const m = Math.ceil(
      d.reduce(
        (acc, cur) => (Number(cur.hours) > acc ? Number(cur.hours) : acc),
        0
      )
    );

    return [d, m];
  }, [projectDurations]);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Projects worked on</CardTitle>
        <CardDescription>Hours spent</CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="min-h-[200px] w-full">
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <YAxis domain={[0, maxY + 2]} dataKey={"hours"} />
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
