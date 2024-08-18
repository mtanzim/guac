import { useEffect, useState } from "react";

import { Banner } from "./Banner";
import { Login } from "./Login";
import { Plot } from "./Plot";
import { DEFAULT_DAY_RANGE, TOKEN_KEY } from "./utils";
import { DateRange } from "react-day-picker";
import { addDays, endOfToday } from "date-fns";
import { ThemeProvider } from "./components/theme-provider";

function App() {
  const [token, setToken] = useState<string | null>(null);
  const isAuthenticated = !!token;

  const [date, setDate] = useState<DateRange | undefined>({
    from: addDays(endOfToday(), -DEFAULT_DAY_RANGE),
    to: endOfToday(),
  });

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
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <Banner date={date} setDate={setDate} onLogout={onLogout} />
      {date?.from && date?.to && (
        <Plot
          start={date?.from}
          end={date?.to}
          onLogout={onLogout}
          token={token}
        />
      )}
    </ThemeProvider>
  );
}

export default App;
