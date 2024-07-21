import { useEffect, useState } from "react";

import { DailyChart } from "./DailyChart";
import { StatsData } from "./data-types";
import { LanguageChart } from "./LanguageChart";
import { Login } from "./Login";
import { ProjectChart } from "./ProjectChart";
import { BASE_URL, DEFAULT_DAY_RANGE, getDateRange, TOKEN_KEY } from "./utils";
import { Banner } from "./Banner";

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
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-4">
      {/* <code>{JSON.stringify(data?.dailyDuration, null, 2)}</code> */}
      {data?.dailyDuration && (
        <DailyChart dailyDuration={data?.dailyDuration} />
      )}
      {data?.languageStats && (
        <LanguageChart languageDurations={data?.languageStats?.durations} />
      )}
      {data?.projectStats && (
        <ProjectChart projectDurations={data?.projectStats?.durations} />
      )}
    </div>
  );
}

export default App;
