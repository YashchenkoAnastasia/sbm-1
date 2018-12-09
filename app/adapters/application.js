import DS from 'ember-data';
import config from '../config/environment';
import DataAdapterMixin from 'ember-simple-auth/mixins/data-adapter-mixin';

export default DS.RESTAdapter.extend(DataAdapterMixin,{
  namespace: 'api',
  host: `${config.protocol}://${config.host}`,
  authorizer: 'authorizer:custom'
});
