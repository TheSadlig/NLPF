package main
import ("github.com/rhinoman/couchdb-go"
	"time"
	"github.com/nu7hatch/gouuid"
)

type ViewResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []interface{} `json:"rows,omitempty"`
}

type MultiReadResponse struct {
	TotalRows int            `json:"total_rows"`
	Offset    int            `json:"offset"`
	Rows      []MultiReadRow `json:"rows"`
}

type MultiReadRow struct {
	Id  string       `json:"id"`
	Key string       `json:"key"`
	Doc interface{} `json:"doc"`
}

type ListResult struct {
	Id  string       `json:"id"`
	Key interface{} `json:"key"`
	//Value string       `json:"value"`
}

type ListResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []ListResult `json:"rows,omitempty"`
}

type View struct {
	Map    string `json:"map"`
	Reduce string `json:"reduce,omitempty"`
}

type DesignDocument struct {
	Language string            `json:"language"`
	Views    map[string]View   `json:"views"`
	Lists    map[string]string `json:"lists"`
}

func getAuth() couchdb.BasicAuth {
	return couchdb.BasicAuth{"admin", "admin"}
}

func getConn () *couchdb.Connection {
	var timeout = time.Duration(50000 * time.Millisecond)
	conn, _ := couchdb.NewConnection("127.0.0.1",5984,timeout)
	return conn

}
func createDB () {
	conn := getConn()
	auth := getAuth()
	conn.CreateDB("nlpf", &auth);
}

func getDB () *couchdb.Database {
	conn := getConn()
	auth := getAuth()
	db := conn.SelectDB("nlpf", &auth)
	
	return db
}

func getUUID() string {
	out, _ := uuid.NewV4()
	return out.String()
}

func createDummyProject() string{
	db := getDB()

	u := getEmptyUser()
	u.Lastname = "Le fossoyeur"
	u.Firstname = "de films"
	u.Mail = "fossoyeur@hotmail.com"
	u.Password = "Coucou"

	db.Save(&u, u.ID, "")
	
	project := getEmptyProject()
	
	project.Name = "Episode du fossoyeur de film"
	project.Description = "« Le Fossoyeur de Films » est le concept principal de ma chaîne Youtube et me permet d’analyser des thématiques pointues sur un ton fantaisiste. La chaîne abrite aussi d’autres émissions : « L’après-séance » (des chroniques de films récemment sortis en salles), « Retour sur le futur » (où il est question de futures sortes et de ce que l’on peut en attendre) ou encore « Dead Watts » (des vidéos musicales décalées). Je réalise et mets en ligne trois vidéos par mois en moyenne, voire quatre ou plus selon la productivité et l’actualité du moment. Je travaille seul sur mes productions et elles me demandent tellement de temps, d’énergie et d’investissement que je ne pourrais pas mener ce projet autrement qu’à 200%. Voulant continuer à vous proposer des contenus réguliers et de qualité, j’ai donc décidé de faire de cette activité de vidéaste mon métier à temps plein. Vivre d’une passion : un pari assez incroyable que vous me permettez déjà de relever grâce à votre soutien et à votre curiosité, ce pour quoi je vous suis extrêmement reconnaissant. Toutefois, à moins de cumuler plusieurs millions de vues par vidéo, il est très difficile de vivre correctement de cette activité. En attendant de grandir, il faut donc se diversifier. "

	project.User_ID= u.ID
	project.Date = "03/12/2015"
	
	db.Save(&project, project.ID, "")

	r:= getEmptyReward()	
	r.Title = "Pour 1€ et plus"
	r.Description = "Vous avez accès à mon fil d’actualité Tipz"
	r.Value = 1
	r.Project_ID = project.ID
	db.Save(&r, r.ID, "")

	r2:= getEmptyReward()	
	r2.Title = "Pour 3€ et plus"
	r2.Description = "Vous avez accès à mon fil d'actualité Tipz et je vous envoie le générique du Fossoyeur de Films en mp3 et/ou .wav"
	r2.Value = 3
	r2.Project_ID = project.ID
	db.Save(&r2, r2.ID, "")
	
	r3:= getEmptyReward()	
	r3.Title = "Pour 5€ et plus"
	r3.Description = " Vous avez accès à mon fil d'actualité Tipz + je vous envoie le générique du Fossoyeur de Films en mp3 et/ou .wav + vous avez accès en avant-première au bêtisier vidéo de l'épisode soutenu "
	r3.Value = 5
	r3.Project_ID = project.ID
	db.Save(&r3, r3.ID, "")
	


	emptyProject := Project{}
	db.Read(project.ID, &emptyProject, nil)

	return project.ID
}

func getUserById(userID string) *User {
	db := getDB();
	
	result := User{}
	_, _ = db.Read(userID, &result, nil)

	return &result
}





func createView(viewName string, designName string, mapCode string, reduceCode string) {
	db := getDB();
	
	
	view := View{
		Map: mapCode}
	if (reduceCode != "") {
		view.Reduce = reduceCode
	}
		
	views := make(map[string]View)
	
	views[viewName] = view
	ddoc := DesignDocument{
		Language: "javascript",
		Views:    views}
	
	db.SaveDesignDoc(designName, ddoc, "")


}
func createViews() {
	createDummyProject()

	createView("get_projects", "projects", "function(doc) {\n  if (doc.Type == \"Project\"){\n    emit(doc.ID, doc)\n  }\n}", "")
	createView("get_rewards", "rewards", "function(doc) {\n if (doc.Type == \"Reward\") {\n    emit(doc.Project_ID, doc);\n  }\n}", "")
	createView("get_users_by_mail", "user_mail", "function(doc) {\n if (doc.Type == \"User\") {\n    emit(doc.Mail, doc);\n  }\n}", "")
	createView("get_users", "user", "function(doc) {\n if (doc.Type == \"User\") {\n    emit(doc.ID, doc);\n  }\n}", "")
	createView("get_investments_number_by_reward",
		"investment",
		"function(doc) {\n if (doc.Type == \"Investment\")\n emit(doc.Reward_ID, 1);\n }",
		"function(keys, values) { \n return sum(values); \n }")
	
}
