import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router } from "react-router-dom";
import { useStrict } from "mobx";
import { Provider } from "mobx-react";
import { sightingsStore } from "./stores";
import App from "./components/App";
import "./index.css";

useStrict(true);

const stores = { sightingsStore };

ReactDOM.render(
  <Provider {...stores}>
    <Router>
      <App />
    </Router>
  </Provider>,
  document.getElementById("root")
);
