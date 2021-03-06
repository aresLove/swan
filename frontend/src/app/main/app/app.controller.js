(function () {
  'use strict';

  angular
    .module('swan')
    .controller('AppController', AppController);

  /** @ngInject */
  function AppController(appBackend, $stateParams) {
    var params = {appId: $stateParams.appId};

    var vm = this;
    vm.app = {};

    activate();

    function activate() {
      getAppInfo()
    }

    function getAppInfo() {
      if (params.appId) {
        appBackend.app(params).get(function (data) {
          vm.app = data;
        });
      }
    }
  }
})();
