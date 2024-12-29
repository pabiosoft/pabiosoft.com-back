package dbchecker

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// RunCreateTables exécute les migrations si `execute` est vrai et liste les tables créées
func RunCreateTables(db *sql.DB, execute bool) {
	if !execute {
		fmt.Println("Les migrations n'ont pas été exécutées.")
		return
	}

	// Log d'information
	log.Println("Démarrage de l'exécution des migrations...")

	// Charger le fichier SQL
	sqlFile := "cmd/migrations/create_tables.sql"
	log.Printf("Chargement du fichier SQL : %s\n", sqlFile)

	content, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("Impossible de lire le fichier SQL : %v", err)
	}

	// Log du contenu du fichier SQL
	log.Println("Contenu du fichier SQL chargé avec succès :")
	log.Println(string(content))

	// Exécuter chaque commande SQL séparément
	queries := strings.Split(string(content), ";")
	for i, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		// Log pour chaque requête
		log.Printf("Exécution de la requête %d : %s\n", i+1, query)

		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Erreur lors de l'exécution de la requête %d : %v\n", i+1, err)
		} else {
			log.Printf("Requête %d exécutée avec succès.\n", i+1)
		}
	}

	// Lister les tables créées
	listTables(db)

	// Fin des migrations
	log.Println("Toutes les migrations ont été exécutées avec succès.")
}

func listTables(db *sql.DB) {
	log.Println("Listing des tables créées dans la base de données...")
	rows, err := db.Query("SELECT tablename FROM pg_tables WHERE schemaname = 'public';")
	if err != nil {
		log.Fatalf("Erreur lors de la récupération des tables : %v", err)
	}
	defer rows.Close()

	log.Println("Tables dans le schéma 'public' :")
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Printf("Erreur lors de la lecture du résultat : %v", err)
			continue
		}
		log.Printf("- %s", tableName)
	}
}

// RunCreateTables exécute les migrations si `execute` est vrai
func RunCreateTablesSimple(db *sql.DB, execute bool) {
	if !execute {
		fmt.Println("Les migrations n'ont pas été exécutées.")
		return
	}

	// Charger le fichier SQL
	sqlFile := "cmd/migrations/create_tables.sql"
	content, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("Impossible de lire le fichier SQL : %v", err)
	}

	// Exécuter le script SQL
	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatalf("Erreur lors de l'exécution du script SQL : %v", err)
	}

	fmt.Println("Migration exécutée avec succès !")
}
