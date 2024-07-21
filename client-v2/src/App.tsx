import { useEffect, useState } from "react";

import { Banner } from "./Banner";
import { Login } from "./Login";
import { TOKEN_KEY } from "./utils";
import { Plot } from "./Plot";

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

export default App;
