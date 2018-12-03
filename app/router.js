import EmberRouter from '@ember/routing/router';
import config from './config/environment';

const Router = EmberRouter.extend({
  location: config.locationType,
  rootURL: config.rootURL
});

Router.map(function() {
  this.route('root', {path: '/'}, function() {
    this.route('members');
    this.route('projects');
    this.route('posts', function() {
      this.route('create');
      this.route('show', {path: '/:id'});
      this.route('edit', {path: '/:id/edit'});
    });
  });
});

export default Router;
