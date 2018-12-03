'use strict';

module.exports = function(environment) {
  let ENV = {
    modulePrefix: 'amir-blog-frontend',
    environment,
    rootURL: '/',
    locationType: 'auto',
    EmberENV: {
      FEATURES: {
        // Here you can enable experimental features on an ember canary build
        // e.g. 'with-controller': true
      },
      EXTEND_PROTOTYPES: {
        // Prevent Ember Data from overriding Date.parse.
        Date: false
      }
    },

    APP: {}
  };

  if (environment === 'development') {
    ENV.protocol = 'http';
    ENV.host = 'localhost:3000';
  }

  if (environment === 'production') {
    ENV.protocol = 'http';
    ENV.host = 'localhost:3000';
  }

  ENV['ember-simple-auth'] = {
    authenticationRoute: 'root',
    routeAfterAuthentication: 'root',
    routeIfAlreadyAuthenticated: 'root',
  };

  ENV['ember-simple-auth-token'] = {
    authorizationHeaderName: 'Authorization',
    authorizationPrefix: ' ',
    tokenPropertyName: 'token',
    serverTokenEndpoint: `${ENV.protocol}://${ENV.host}/api/users/authenticate`
  };

  return ENV;
};
