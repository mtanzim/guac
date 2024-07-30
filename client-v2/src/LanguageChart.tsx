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
import { useMemo } from "react";
import { SkeletonChart } from "./SkeletonChart";

export function LanguageChart({
  languageDurations,
  loading,
}: {
  languageDurations?: StatsData["languageStats"]["durations"];
  loading: boolean;
}) {
  const chartConfig = {
    hours: {
      label: "Hours",
    },
  } satisfies ChartConfig;

  const [chartData, maxY] = useMemo(() => {
    const d = (languageDurations || [])
      .slice()
      .sort((a, b) => b.minutes - a.minutes)
      .slice(0, TOP_N_LANGUAGES)
      .map((d) => ({
        hours: (d.minutes / 60).toFixed(2),
        language: d.language,
      }));

    const restMinutes = (languageDurations || [])
      .slice()
      .sort((a, b) => b.minutes - a.minutes)
      .slice(TOP_N_LANGUAGES)
      .reduce((acc, cur) => acc + cur.minutes, 0);

    d.push({
      hours: (restMinutes / 60).toFixed(2),
      language: `Rest`,
    });
    const m = Math.ceil(
      d.reduce(
        (acc, cur) => (Number(cur.hours) > acc ? Number(cur.hours) : acc),
        0
      )
    );
    return [d, m];
  }, [languageDurations]);
  if (loading) {
    return <SkeletonChart />;
  }
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
            <YAxis domain={[0, maxY + 2]} dataKey={"hours"} />
            <XAxis
              dataKey="language"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
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
