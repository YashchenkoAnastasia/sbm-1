import AjaxService from 'ember-ajax/services/ajax';
import config from '../config/environment';
import { computed } from '@ember/object';
import {inject as service} from '@ember/service';

export default AjaxService.extend({
  session: service(),
  host: `${config.protocol}://${config.host}`,
  contentType: 'application/json; charset=utf-8',
  headers: computed('session.data.authenticated.token', {
    get() {
      let headers = {};
      const token = this.get('session.data.authenticated.token');
      if (token) {
        headers['Authorization'] = `Bearer ${token}"`;
      }
      return headers;
    }
  })
});
