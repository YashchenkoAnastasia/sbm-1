import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | root/members', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:root/members');
    assert.ok(route);
  });
});
