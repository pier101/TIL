#include "mainGame.h"
#include "enemy.h"
#include "player.h"

int main() 
{
	//mainGame mg;
	//mainGame mg1(3);
	//mainGame mg2(100, 50);

	player* pl = new player;
	enemy* em = new enemy;

	pl->printPlayer();
	
	em->printEnemy();
	//���ʹ� Ÿ�� ���� �Ϳ� ���� HP�缳��
	em->setEnemyHP(em->getEnemyHP() - pl->getPlayerAtk());

	em->printEnemy();

	delete pl;
	delete em;
	return 0;
}