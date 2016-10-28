tipz.controller('tipzHandler', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Inscription";

    $scope.InscriptionMail = "";
    $scope.InscriptionPassword = "";
    $scope.InscriptionLastname = "";
    $scope.InscriptionFirstname = "";
    
    $scope.inscriptionClick = function() {
        if ($scope.InscriptionMail == "" || $scope.InscriptionMail.indexOf('@') < 0) // no @ in mail
            $scope.$parent.$parent.error = "Email incorrect";
        else if ($scope.InscriptionLastname == "")
            $scope.$parent.$parent.error = "Nom incorrect";
        else if ($scope.InscriptionFirstname == "")
            $scope.$parent.$parent.error = "Prenom incorrect";
        else if ($scope.InscriptionPassword == "")
            $scope.$parent.$parent.error = "Mot de passe incorrect";
        else
        {
    	    $http({
                method: 'POST',
                url: "http://localhost:9090/api/createUser",
                data: {data:
                       { data:
                         {lastname: $scope.InscriptionLastname,
                          firstname: $scope.InscriptionFirstname,
                          mail:  $scope.InscriptionMail,
                          password: $scope.InscriptionPassword,
                         }
                       }
                      },
                transformRequest: function(obj) {
                    var str = [];
                    for(var p in obj)
                        str.push(encodeURIComponent(p) + "=" + encodeURIComponent(JSON.stringify(obj[p]) ) );
                    return str.join("&");
                },
                headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
            }).success(function(data){
                console.log(data)
                if (data.success == true)
		    $scope.$parent.$parent.changePage("connection.html");
                else
                    $scope.$parent.$parent.error = "Cet email est deja pris";
            })
        }
    }
}]);
