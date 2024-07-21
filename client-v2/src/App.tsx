import { Button } from "@/components/ui/button";

import { Input } from "@/components/ui/input";
import { MouseEventHandler, useEffect, useState } from "react";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";

import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { StatsData } from "./data-types";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./components/ui/card";

const BASE_URL = "http://localhost:8080";

function Login({ onLogin }: { onLogin: (t: string) => void }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [errMsg, setErrMsg] = useState<string | null>(null);
  const handleSubmit: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault();
    setLoading(true);
    setErrMsg(null);

    try {
      const res = await fetch(`${BASE_URL}/api/v1/login`, {
        method: "POST",
        body: JSON.stringify({
          username,
          password,
        }),
      });
      const body = await res.json();
      if (res.ok) {
        onLogin(body?.token);
        return;
      }
      throw new Error("failed to login");
    } catch (err) {
      console.log(err);
      setErrMsg("Something went wrong");
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className="w-full max-w-sm items-center">
      <Input
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        className="mt-4"
        type="username"
        placeholder="Username"
      />
      <Input
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        className="mt-4"
        type="password"
        placeholder="Password"
      />
      {errMsg && <p className="text-red-500">{errMsg}</p>}
      {isLoading && <p className="text-slate-400 animate-pulse">Loading...</p>}
      <Button
        className="mt-4 w-1/4 float-end"
        type="submit"
        onClick={handleSubmit}
      >
        Login
      </Button>
    </div>
  );
}

const TOKEN_KEY = "WakaToken";

function App() {
  const [token, setToken] = useState<string | null>(null);
  const isAuthenticated = !!token;

  useEffect(() => {
    const curToken = window.localStorage.getItem(TOKEN_KEY);
    if (curToken) {
      setToken(curToken);
    }
  }, []);

  const onLogin = (token: string) => {
    window.localStorage.setItem(TOKEN_KEY, token);
    setToken(token);
  };
  const onLogout = () => {
    window.localStorage.removeItem(TOKEN_KEY);
    setToken(null);
  };

  if (!isAuthenticated) {
    return (
      <div className="mt-64 flex flex-col justify-center items-center">
        <h2 className="text-xl">Login to Guac Dashboard</h2>
        <Login onLogin={onLogin} />
      </div>
    );
  }

  return (
    <>
      <Banner onLogout={onLogout} />
      <Plot onLogout={onLogout} token={token} />
    </>
  );
}

const DEFAULT_DAY_RANGE = 7;

function getDateRange(days: number) {
  const formatDateForReq = (date: Date) => {
    return date.toISOString().split("T")[0];
  };

  const endDate = new Date();
  endDate.setDate(endDate.getDate() - 1);
  const ending = formatDateForReq(endDate);
  const startDate = new Date();
  startDate.setDate(startDate.getDate() - days);
  const starting = formatDateForReq(startDate);
  return { starting, ending };
}

function Plot({ onLogout, token }: { onLogout: () => void; token: string }) {
  const [data, setData] = useState<null | StatsData>(null);
  const [loading, setLoading] = useState(false);
  const [errMsg, setErrMsg] = useState<null | string>(null);
  useEffect(() => {
    const { starting, ending } = getDateRange(DEFAULT_DAY_RANGE);
    const url =
      starting && ending
        ? `${BASE_URL}/api/v1/data?start=${starting}&end=${ending}`
        : `${BASE_URL}/api/v1/data`;
    setLoading(true);
    fetch(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        }
        throw new Error("cannot fetch");
      })
      .then((d) => setData(d))
      .catch((err) => {
        console.log(err);
        setErrMsg("Something went wrong");
        onLogout();
      })
      .finally(() => setLoading(false));
  }, [token, onLogout]);

  if (loading) {
    return <p>Loading...</p>;
  }
  if (errMsg) {
    return <p>{errMsg}</p>;
  }
  return (
    <>
      {/* <code>{JSON.stringify(data?.dailyDuration, null, 2)}</code> */}
      {data?.dailyDuration && (
        <DailyChart dailyDuration={data?.dailyDuration} />
      )}
    </>
  );
}

const toCustomDateStr = (d: string): string => {
  return new Date(d).toLocaleDateString(undefined, {
    month: "short",
    day: "2-digit",
  });
};

function DailyChart({
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

  const start = toCustomDateStr(dailyDuration.at(0)?.date || "");
  const end = toCustomDateStr(dailyDuration.at(-1)?.date || "");

  return (
    <Card>
      <CardHeader>
        <CardTitle>Daily time spent coding</CardTitle>
        <CardDescription>
          {start} - {end}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="min-h-[200px] w-full">
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <YAxis dataKey={"hours"}/>
            <XAxis
              dataKey="date"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={toCustomDateStr}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Bar dataKey={"hours"} radius={4} />;
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}

function Banner({ onLogout }: { onLogout: () => void }) {
  return (
    <div className="flex bg-slate-400 p-6 gap-8 items-center">
      <h2 className="text-xl ">Waka Dashboard</h2>
      <Button className="mr-4" onClick={onLogout}>
        Logout
      </Button>
    </div>
  );
}

export default App;
