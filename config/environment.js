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
    // here you can enable a production-specific feature
  }

  return ENV;
};
