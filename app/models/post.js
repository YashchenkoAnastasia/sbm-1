import DS from 'ember-data';
import { computed } from '@ember/object';
import { validator, buildValidations } from 'ember-cp-validations';

const Validations = buildValidations({
  title:  [
    validator('presence', true)
  ],
  body:  [
    validator('presence', true)
  ]
});

export default DS.Model.extend(Validations, {
  title: DS.attr('string'),
  body:  DS.attr('string'),

  shortBody: computed('body', function () {
    return this.body.substring(0, 70);
  })
});
