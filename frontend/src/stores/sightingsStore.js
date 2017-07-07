import { extendObservable, action, runInAction } from "mobx";
import * as api from "./api";

class SightingsStore {
  constructor() {
    extendObservable(this, {
      page: null,
      selected: null,
      loading: false,

      startLoading: action(() => {
        this.loading = true;
        this.error = null;
      }),

      setPage: action((page, error) => {
        this.loading = false;
        this.page = error ? null : page;
        this.error = error;
      }),

      fetchAll: async search => {
        this.startLoading();

        let response, error;

        try {
          response = await api.getSightings(search);
        } catch (err) {
          error = err;
        }

        this.setPage(response, error);
      },

      fetchPage: async url => {
        this.startLoading();

        let response, error;

        try {
          response = await api.getSightingsPage(url);
        } catch (err) {
          error = err;
        }

        this.setPage(response, error);
      },
      fetchOne: async id => {
        this.startLoading();

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
