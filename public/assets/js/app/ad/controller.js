/**
 * Created by clude on 9/24/15.
 */
ni.NiNgModule
    .controller('AdsCtrl', ['$scope', '$http', function ($scope, $http) {
        var vm = $scope.vm = {};
        var fn = $scope.fn = {};

        fn.reloadFrame = function(){
            vm.frameSrc = '/ads/creative?rnd=' + Math.random()* 100;
        }

        fn.getCachedCreative = function(){
            var angularHttpSetting = {
                method:     'GET',
                url:        '/api/creative/get',
                data:       {},
                headers:    {}
            }

            $http(angularHttpSetting).
                success(function(data, status, headers, config) {
                    vm.creative = data.creative;
                }).
                error(function(data, status, headers, config) {
                    console.log(data);
                });
        }

        fn.postCreative = function(){
            var angularHttpSetting = {
                method:     'POST',
                url:        '/api/creative/save',
                data:       {creative: vm.creative},
                headers:    {}
            }

            $http(angularHttpSetting).
                success(function(data, status, headers, config) {
                    fn.reloadFrame();
                }).
                error(function(data, status, headers, config) {
                    console.log(data);
                });
        }

        fn.getCachedCreative();
        fn.reloadFrame();

    }])

;