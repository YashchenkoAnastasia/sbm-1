import Controller from '@ember/controller';

export default Controller.extend({
  actions: {
    authenticate() {
      this.set('errorMessage', null);
      const credentials = this.getProperties('login', 'password'), authenticator = 'authenticator:jwt';
      this.get('session').authenticate(authenticator, credentials).then(() => {
        this.toggleProperty('modal');
      }).catch((reason) => {
        this.set('errorMessage', reason.payload.error);
      });
    }
  }
});
