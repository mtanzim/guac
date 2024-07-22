import * as React from "react";
import { Cell, Label, Pie, PieChart, Sector } from "recharts";
import { PieSectorDataItem } from "recharts/types/polar/Pie";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartStyle,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { StatsData } from "./data-types";
import { getColor, TOP_N_LANGUAGES } from "./utils";

const chartConfig = {
  language: {
    label: "Language",
  },
} satisfies ChartConfig;

export function LanguagePct({
  rawPercentages,
}: {
  rawPercentages: StatsData["languageStats"]["percentages"];
}) {
  const percentages = rawPercentages
    .slice()
    .sort((a, b) => b.percentage - a.percentage)
    .slice(0, TOP_N_LANGUAGES + 1);
  const totalPct = percentages.reduce((acc, cur) => acc + cur.percentage, 0);
  percentages.push({ percentage: 100 - totalPct, language: "Rest" });

  const id = "pie-interactive";
  const [active, setActive] = React.useState(percentages[0].language);

  const activeIndex = React.useMemo(
    () => percentages.findIndex((item) => item.language === active),
    [active, percentages]
  );
  const langs = React.useMemo(
    () => percentages.map((item) => item.language),
    [percentages]
  );

  return (
    <Card data-chart={id} className="flex flex-col">
      <ChartStyle id={id} config={chartConfig} />
      <CardHeader className="flex-row items-start space-y-0 pb-0">
        <div className="grid gap-1">
          <CardTitle>Languages Used</CardTitle>
          <CardDescription>Percentage</CardDescription>
        </div>
        <Select value={active} onValueChange={setActive}>
          <SelectTrigger
            className="ml-auto h-7 w-[130px] rounded-lg pl-2.5"
            aria-label="Select a value"
          >
            <SelectValue placeholder="Select language" />
          </SelectTrigger>
          <SelectContent align="end" className="rounded-xl">
            {langs.map((key) => {
              return (
                <SelectItem
                  key={key}
                  value={key}
                  className="rounded-lg [&_span]:flex"
                >
                  <div className="flex items-center gap-2 text-xs">
                    <span
                      className="flex h-3 w-3 shrink-0 rounded-sm"
                      style={{
                        backgroundColor: getColor(key),
                      }}
                    />
                    {key}
                  </div>
                </SelectItem>
              );
            })}
          </SelectContent>
        </Select>
      </CardHeader>
      <CardContent className="flex flex-1 justify-center pb-0">
        <ChartContainer
          id={id}
          config={chartConfig}
          className="mx-auto aspect-square w-full max-w-[800px]"
        >
          <PieChart>
            <ChartTooltip content={<ChartTooltipContent />} />
            <Pie
              data={percentages}
              dataKey="percentage"
              nameKey="language"
              innerRadius={60}
              strokeWidth={5}
              activeIndex={activeIndex}
              activeShape={({
                outerRadius = 0,
                ...props
              }: PieSectorDataItem) => (
                <g>
                  <Sector {...props} outerRadius={outerRadius + 10} />
                  <Sector
                    {...props}
                    outerRadius={outerRadius + 25}
                    innerRadius={outerRadius + 12}
                  />
                </g>
              )}
            >
              {percentages.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={getColor(entry.language)} />
              ))}
              <Label
                content={({ viewBox }) => {
                  if (viewBox && "cx" in viewBox && "cy" in viewBox) {
                    return (
                      <text
                        x={viewBox.cx}
                        y={viewBox.cy}
                        textAnchor="middle"
                        dominantBaseline="middle"
                      >
                        <tspan
                          x={viewBox.cx}
                          y={viewBox.cy}
                          className="fill-foreground text-xl font-bold"
                        >
                          {active}
                        </tspan>
                        <tspan
                          x={viewBox.cx}
                          y={(viewBox.cy || 0) + 24}
                          className="fill-muted-foreground"
                        >
                          {Number(
                            percentages.find((p) => p.language === active)
                              ?.percentage || 0
                          ).toFixed(1)}
                          %
                        </tspan>
                      </text>
                    );
                  }
                }}
              />
            </Pie>
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
