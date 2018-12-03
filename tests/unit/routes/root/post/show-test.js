import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | root/post/show', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:root/post/show');
    assert.ok(route);
  });
});
