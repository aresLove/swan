(function () {
  'use strict';

  angular
    .module('swan')
    .config(routerConfig);

  /** @ngInject */
  function routerConfig($stateProvider, $urlRouterProvider) {
    $stateProvider
      .state('home', {
        templateUrl: 'app/main/main.html',
        controller: 'MainController',
        controllerAs: 'vm'
      })
      .state('home.cluster', {
        url: '/cluster',
        templateUrl: 'app/main/cluster/cluster.html',
        controller: 'ClusterController',
        controllerAs: 'vm'
      })
      .state('home.user', {
        url: '/user?fields',
        templateUrl: 'app/main/user/user.html',
        controller: 'UserController',
        controllerAs: 'vm'
      })
      .state('home.app', {
        url: '/app?appId',
        templateUrl: 'app/main/app/app.html',
        controller: 'AppController',
        controllerAs: 'vm'
      })
      .state('home.task', {
        url: '/task?appId&taskIndex',
        templateUrl: 'app/main/task/task.html',
        controller: 'TaskController',
        controllerAs: 'vm'
      });

    $urlRouterProvider.otherwise(function ($injector) {
      var $state = $injector.get('$state');
      $state.go('home.cluster');
    });
  }

})();
