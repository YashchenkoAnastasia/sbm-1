import DS from 'ember-data';
import config from '../config/environment';

export default DS.RESTAdapter.extend({
  namespace: 'api/1',
  host: `${config.protocol}://${config.host}`
});
