import { camelizeKeys } from "humps";

class Http {
  async fetch(url) {
    const req = new window.Request(url, {
      method: "GET",
      headers: new window.Headers({
        "Content-Type": "application/json"
      })
    });

    const response = await window.fetch(req);
    const payload = await response.json();

    return camelizeKeys(payload);
  }
}

export default Http;
