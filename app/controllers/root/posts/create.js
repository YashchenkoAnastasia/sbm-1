import Controller from '@ember/controller';
import { task }   from 'ember-concurrency';

export default Controller.extend({
  save: task(function * () {
    yield this.get('model').save().then((model) => {
      this.get('notifications').success('Вы успешно добавили пост', {
        autoClear: true,
        clearDuration: 10000
      });
      this.transitionToRoute('root.posts.show', model.id);
    }).catch((data) => {
      this.get('notifications').error('Что-то пошло не так.', {
        autoClear: true,
        clearDuration: 10000
      });
    });
  })

});
