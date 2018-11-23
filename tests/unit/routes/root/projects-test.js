import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | root/projects', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:root/projects');
    assert.ok(route);
  });
});
