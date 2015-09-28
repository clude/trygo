/**
 * Created by clude on 9/22/15.
 */
ni.NiNgModule
    .controller('DmMainCtrl', ['$scope', 'NiHttp', function ($scope, NiHttp) {
        var gvm = $scope.gvm = {};
        var gfn = $scope.gfn = {};

        gvm.aerospkie = {};
        gvm.aerospkie.namespace = 'camel';

        gvm.tabs = {
            ads: {active: false},
            de: {active: false},
            dsp: {active: false}
        }

        var loc = window.location.href;
        if(loc.indexOf('/de') >=0){
            gvm.tabs.de.active = true;
        }else if(loc.indexOf('/dsp') >= 0){
            gvm.tabs.dsp.active = true;
        }else{
            gvm.tabs.ads.active = true;
        }

        gvm.isInit = false;
        gfn.RedirectTo = function(target){
            if(gvm.isInit){
                gvm.isUnloaded = true;
                window.location.href = target
            }
            gvm.isInit = true;
        }
    }])
    .controller('ServersCtrl', ['$scope', 'Api', function ($scope, Api) {
        var vm = $scope.vm = {};
        var fn = $scope.fn = {};

        fn.getServerStatus = function(){
            Api.getRecord($scope.gvm.tabs.de.active? 'de_ctrl' : 'dsp_ctrl',  'servers')
                .success(function(rst){
                    vm.servers = rst.data['bin']; // pretty print json
                })
                .error(function(rst){
                    vm.servers = {};
                })
        }

        fn.getServerStatus();
    }])
    .controller('RecordCtrl', ['$scope', 'Api', function ($scope, Api) {
        var vm = $scope.vm = {};
        var fn = $scope.fn = {};

        vm.results=""

        vm.param = {}

        fn.getRecord = function(){
            Api.getRecord(vm.param.set, vm.param.key)
                .success(function(rst){
                    vm.results = JSON.stringify(rst.data, null, 2); // pretty print json
                })
                .error(function(rst){
                    vm.results = rst.msg;
                })
        }
    }])
    .controller('DeUpdaterCtrl', ['$scope', 'Api', '$filter', function ($scope, Api, $filter) {
        var vm = $scope.vm = {};
        var fn = $scope.fn = {};

        vm.param = {};

        vm.dt = new Date();

        fn.getRecord = function(){
            var set = 'de_ctrl';
            var dt = $filter('date')(vm.dt, 'yyyyMMdd');
            var key = String.format("{0}_ctrl_{1}_{2}", "s", dt, vm.param.id);

            Api.getRecord(set, key)
                .success(function(rst){
                    vm.results = rst.data['bin'];
                })
                .error(function(rst){
                    vm.results = rst.msg;
                })
        }
    }])
    .controller('DspUpdaterCtrl', ['$scope', 'Api', '$filter', function ($scope, Api, $filter) {
        var vm = $scope.vm = {};
        var fn = $scope.fn = {};

        vm.param = {}

        vm.types = [
            {key:"s", name:'投放'},
            {key:"c", name:'订单'},
            {key:"a", name:'广告主'}
        ]

        vm.dt = new Date();

        vm.results = {}

        fn.getAll = function(){
            var day = $filter('date')(vm.dt, 'yyyyMMdd');
            if(vm.param.sid){
                fn.getRecord('s', 'global', vm.param.sid);
                fn.getRecord('s', day, vm.param.sid);
            }
            if(vm.param.cid){
                fn.getRecord('c', 'global', vm.param.cid);
                fn.getRecord('c', day, vm.param.cid);
            }
            if(vm.param.aid){
                fn.getRecord('a', 'global', vm.param.aid);
                fn.getRecord('a', day, vm.param.aid);
            }
        }

        fn.getRecord = function(type, day, id){

            var set = 'dsp_ctrl';
            var key = String.format("{0}_ctrl_{1}_{2}", type, day, id);

            vm.results[type] = vm.results[type] || {};

            vm.results[type].day = {
                exp: 100,
                show: 10,
                click: 5,
                exp_127: 1
            }

            vm.results[type].global = {
                exp: 77,
                show: 8,
                click: 5,
                exp_127: 1
            }
            return;
            Api.getRecord(set, key)
                .success(function(rst){
                    if(day){
                        vm.results[type].day = rst.data['bin'] || {};
                    }else{
                        vm.results[type].global = rst.data['bin'] || {};
                    }
                })
                .error(function(rst){
                    if(day){
                        vm.results[type].day = {};
                    }else{
                        vm.results[type].global = {};
                    }
                })
        }
    }])
    .controller('FrequencyCtrl', ['$scope', 'Api', '$filter', function ($scope, Api, $filter) {
        var vm = $scope.vm = {};
        var fn = $scope.fn = {};

        vm.param = {}

        vm.types = [
            {key:"s", name:'投放'},
            {key:"c", name:'订单(Campaign)'},
            {key:"o", name:'订单(Order)'}
        ];
        vm.results = {}

        fn.getAll = function(){
            if(vm.param.sid){
                fn.getRecord('s', vm.param.cookie, vm.param.sid);
            }
            if(vm.param.cid){
                fn.getRecord('c', vm.param.cookie,vm.param.cid);
            }
            if(vm.param.oid){
                fn.getRecord('o', vm.param.cookie,vm.param.oid);
            }

            if(!_.isEmpty(vm.param.cookie)){
                fn.getTag(vm.param.cookie);
            }
        }

        fn.getRecord = function(type, cookie, id){

            var set = String.format("{0}_freq_{1}", $scope.gvm.tabs.de.active? 'de' : 'dsp', type);
            var key = String.format("{0}_{1}", cookie, id);

            vm.results[type] = vm.results[type] || {};
            Api.getRecord(set, key)
                .success(function(rst){
                    vm.results[type] = rst.data;
                })
                .error(function(rst){
                    vm.results[type] = {}
                })
        }

        fn.getTag = function(cookie){
            var set = 'dsp_tag_r';
            var key = cookie;

            vm.results.tag = vm.results.tag || {};
            Api.getRecord(set, key)
                .success(function(rst){
                    vm.results.tag = JSON.stringify(rst.data['bin'], null, 2) ;
                })
                .error(function(rst){
                    vm.results.tag = {}
                })

        }
    }])

;