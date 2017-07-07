import { extendObservable, runInAction } from "mobx";
import * as api from "./api";

class SightingsStore {
  constructor() {
    extendObservable(this, {
      page: null,
      selected: null,
      loading: false,

      fetchAll: async search => {
        runInAction(() => {
          this.loading = true;
          this.error = null;
        });

        let response, error;

        try {
          response = await api.getSightings(search);
        } catch (err) {
          error = err;
        }

        runInAction(() => {
          this.page = response;
          this.error = error;
          this.loading = false;
        });
      },

      fetchPage: async url => {
        runInAction(() => {
          this.loading = true;
          this.error = null;
        });

        let response, error;

        try {
          response = await api.getSightingsPage(url);
        } catch (err) {
          error = err;
        }

        runInAction(() => {
          this.page = response;
          this.error = error;
          this.loading = false;
        });
      },
      fetchOne: async id => {
        runInAction(() => {
          this.loading = true;
          this.error = null;
          this.selected = null;
        });

        let response, error;

        try {
          response = await api.getSighting(id);
        } catch (err) {
          error = err;
        }

        runInAction(() => {
          this.selected = response;
          this.error = error;
          this.loading = false;
        });
      }
    });
  }
}

export default new SightingsStore();
