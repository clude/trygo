<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
    <script src='/static/assets/js/libs/angular.js'></script>
    <script src='/static/assets/js/libs/angular-route.js'></script>
    <script src='/static/assets/js/libs/angular-sanitize.js'></script>
    <script src='/static/assets/js/libs/angular-resource.js'></script>
    <link rel='stylesheet' href='/static/assets/css/bootstrap.css' />
    <link rel='stylesheet' href='/static/assets/css/style.css' />

    <script type="text/javascript">
        angular.module('app', [])
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

    </script>
</head>

<body  ng-app="app"  ng-controller="AdsCtrl">
    <h3>提交创意代码</h3>
    <textarea size="100" ng-model="vm.creative" style="width: 600px; height: 120px" placeholder="输入创意代码"></textarea>
    <input type="button" class="btn" value="提交预览" ng-click="fn.postCreative()" />


    <h3>效果展示 <input type="button" class="btn" value="刷新" ng-click="fn.reloadFrame()" /></h3>
    <div style="border: 2px dotted skyblue; width: 90%; height: 300px; padding: 10px; overflow: auto;">
        <iframe ng-src="{{vm.frameSrc}}" style="width: 100%; height: 100%; border: none;"></iframe>
    </div>
</body>
</html>