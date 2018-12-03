import Base from 'ember-simple-auth/authenticators/base';
import config from '../config/environment';
import { Promise } from 'rsvp';
import { run } from '@ember/runloop';
import { isEmpty } from '@ember/utils';
import {inject as service} from '@ember/service';

export default Base.extend({
  ajax: service(),
  tokenEndpoint: `${config.protocol}://${config.host}/api/users/authenticate`,

  restore(data) {
    return new Promise((resolve, reject) => {
      if (!isEmpty(data.token)) {
        resolve(data);
      } else {
        reject();
      }
    });
  },

  authenticate(creds) {
    const { login, password } = creds;
    const ajax = this.get('ajax');
    const data = {
      login: login,
      password: password
    };

    return new Promise((resolve, reject) => {
      ajax.request(this.tokenEndpoint, {
        method: 'POST',
        data: data
      }).then((response) => {
        run(() => {
          resolve({ token: response.jwt });
        });
      }, (error) => {
        run(() => {
          reject(error);
        });
      });
    });
  },

  invalidate(data) {
    return Promise.resolve(data);
  }
});
