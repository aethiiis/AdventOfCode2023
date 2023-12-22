# AdventOfCode2023
This repository is for the Advent of Code 2023!

Semaine 1\
Jour 1\
Aucun problème en particulier, le problème est bruteforçable. Découverte des runes et du module strings.

Jour 2\
Création de la fonction ReadLines, afin de raccourcir le code pour les prochains jours.

Jour 3\
Premier pic de difficulté pour moi, le code est très long pour traiter tous les cas particuliers. Je pense que la longueur du code pourrait être raccourcie en utilisant une structure du type map[[2]int{}]bool, ce qui sera fait pour les prochains jours.

Semaine 2\
Jour 4\
Le problème est bruteforçable, aucun problème en particulier.

Jour 5\
Gros pic de difficulté, le bruteforce n'est pas envisageable pour la partie 2. Manipulation d'intervalles un peu compliquée au début, mais je suis très satisfait de ma solution qui tourne rapidement.

Jour 6\
Extrêmement simple, résolution d'une équation du second degré.

Jour 7\
Pas de grosse difficulté, mais bon exercice pour l'implémentation de structures.

Jour 8\
Problème assez confusant, repose sur une hypothèse que les données bouclent. 

Jour 9\
Problème facile à résoudre, pourvu qu'on soit au clair sur les indices à regarder.

Jour 10\
Partie 1 facile en suivant le tuyau des deux côtés du point de départ. Partie 2 plus compliquée, car difficilement bruteforçable. J'ai opté pour la solution d'inondation même si compter le nombre de tuyaux du circuit principal est sans doute plus rapide.

Semaine 3\
Jour 11\
Partie 1 simple. Partie 2 un peu plus dure, j'ai opté pour le fait de compter le nombre de lignes et de colonnes étendues entre chaque point, ce qui je pense, est la solution principale.

Jour 12\
Gros pic de difficulté pour moi, car j'ai déjà beaucoup de mal avec la combinatoire. La partie 2 a été élaborée plusieurs jours après le jour 12 et en ayant vu des solutions sur Internet. J'en ai donc profité pour expérimenter avec les channels en Go.

Jour 13\
Problème plus facile que celui de la veille, débugger les fonctions de symétrie fut un peu long.

Jour 14\
Fondamentalement rien de compliqué, il faut juste essayer jusqu'à tomber sur le bon résultat pour trouver l'indice du dernier cycle.

Jour 15\
Partie 1 très simple. Partie 2 un peu plus dure, mais très abordable pourvu que l'on définisse les bonnes structures de données

Jour 16\
Beaucoup de mal à le faire faire fonctionner sur Go alors que facile sur Python. A été beaucoup plus simple une fois les structures appropriées définies. Pour la partie 2, j'ai opté pour du bruteforce, car je ne voyais pas d'autres solutions, mais le code est par conséquent un peu long.

Jour 17\
Problème très similaire à des problèmes des jours précédents, rine de particulier.

Semaine 4\
Jour 18\
Problème simple une fois que l'on connait et réussit à implémenter la formule du lacet.

Jour 19\
Problème relativement simple. Malheureusement, j'avais mal lu la consigne et pensais que l'intervalle commençait à 0 alors qu'il commence en réalité à 0. Cette regrettable confusion m'a coûté deux heures :/.

Jour 20\
Problème un peu plus complexe et qui demandait de bien examiner l'input. Les nombres des cycles que je multiplie pour obtenir le résultat demandé sont premiers donc pas besoin de diviser par le PGCD.

Jour 21\
Première partie du problème similaire à des algorithmes déjà étudié. La seconde partie n'a pas été traitée par manque d'intérêt pour ce genre de problème géométrique et de temps. Je pense savoir comment faire : après avoir examiné l'input, on remarque que le point de départ est au centre du tableau, et que ce-dernier est situé sur une ligne et sur une colonne ne contenant aucun obstacle. Il y a donc des "autoroutes" qui relient le point de départ à tous les blocs, dont la distance de Manhattan est facile à calculer. Il faut ensuite calculer le reste du chemin à parcourir pour chaque bloc, puis gérer les cas à la frontière des blocs.

Jour 22\
Problème non-traité par manque de temps (forum des écoles dans mon ancienne prépa)

Jour 23


Jour 24


Jour 25