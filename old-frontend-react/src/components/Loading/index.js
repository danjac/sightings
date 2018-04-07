import React from "react";

import { Jumbotron } from "reactstrap";

import loadingImg from "./loading.gif";

export default () => {
  return (
    <Jumbotron style={{ textAlign: "center" }}>
      <h1>Loading...</h1>
      <img src={loadingImg} alt="loading" />
    </Jumbotron>
  );
};
