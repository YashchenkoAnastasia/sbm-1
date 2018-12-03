import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | root/posts/create', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:root/posts/create');
    assert.ok(route);
  });
});
