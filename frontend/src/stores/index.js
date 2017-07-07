import Http from "./http";
import SightingsStore from "./sightingsStore";

const http = new Http();

const sightingsStore = new SightingsStore(http);

export { sightingsStore };
