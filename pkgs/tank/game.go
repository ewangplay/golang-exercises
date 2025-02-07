package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Tank 表示坦克
type Tank struct {
	x, y             float64
	direction        int // 0: 上, 1: 右, 2: 下, 3: 左
	directionChanged bool
	hasShot          bool
}

// Bullet 表示子弹
type Bullet struct {
	x, y      float64
	direction int
}

// Wall 表示墙壁
type Wall struct {
	x, y          float64
	width, height float64
	health        int
}

// Game 表示游戏状态
type Game struct {
	playerTank    *Tank
	playerBullets []Bullet
	enemyTanks    []Tank
	enemyBullets  []Bullet
	walls         []Wall
}

// NewGame 创建一个新的游戏实例
func NewGame() *Game {
	return &Game{
		playerTank: &Tank{
			x:         screenWidth / 2,
			y:         screenHeight / 2,
			direction: 0,
		},
		playerBullets: []Bullet{},
		enemyTanks: []Tank{
			{
				x:         100,
				y:         100,
				direction: 2,
			},
		},
		enemyBullets: []Bullet{},
		walls: []Wall{
			{x: 150, y: 150, width: 100, height: 10, health: 3},
			{x: 250, y: 250, width: 150, height: 10, health: 3},
			{x: 400, y: 50, width: 10, height: 100, health: 3},
			{x: 350, y: 300, width: 10, height: 50, health: 3},
		},
	}
}

// Update 更新游戏状态
func (g *Game) Update() error {
	if g.playerTank != nil {
		// 处理坦克移动
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			g.playerTank.direction = 0
			if g.playerTank.y > 0 {
				g.playerTank.y -= tankSpeed
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			g.playerTank.direction = 1
			if g.playerTank.x < screenWidth-20 {
				g.playerTank.x += tankSpeed
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			g.playerTank.direction = 2
			if g.playerTank.y < screenHeight-20 {
				g.playerTank.y += tankSpeed
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			g.playerTank.direction = 3
			if g.playerTank.x > 0 {
				g.playerTank.x -= tankSpeed
			}
		}

		// 处理射击
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			bullet := Bullet{
				x:         g.playerTank.x + 8,
				y:         g.playerTank.y + 8,
				direction: g.playerTank.direction,
			}
			g.playerBullets = append(g.playerBullets, bullet)
		}

		// 更新子弹位置
		for i := 0; i < len(g.playerBullets); i++ {
			switch g.playerBullets[i].direction {
			case 0:
				g.playerBullets[i].y -= bulletSpeed
			case 1:
				g.playerBullets[i].x += bulletSpeed
			case 2:
				g.playerBullets[i].y += bulletSpeed
			case 3:
				g.playerBullets[i].x -= bulletSpeed
			}

			// 检测玩家子弹与敌方坦克的碰撞
			for j := 0; j < len(g.enemyTanks); j++ {
				if checkCollision(g.playerBullets[i].x, g.playerBullets[i].y, 5, 5, g.enemyTanks[j].x, g.enemyTanks[j].y, 20, 20) {
					// 移除敌方坦克
					g.enemyTanks = append(g.enemyTanks[:j], g.enemyTanks[j+1:]...)
					// 移除子弹
					g.playerBullets = append(g.playerBullets[:i], g.playerBullets[i+1:]...)
					i--
					break
				}
			}

			// 检测玩家子弹与墙的碰撞
			for j := 0; j < len(g.walls); j++ {
				if i > 0 && i < len(g.playerBullets) {
					if checkCollision(g.playerBullets[i].x, g.playerBullets[i].y, 5, 5, g.walls[j].x, g.walls[j].y, g.walls[j].width, g.walls[j].height) {
						g.walls[j].health--
						if g.walls[j].health <= 0 {
							g.walls = append(g.walls[:j], g.walls[j+1:]...)
						}
						g.playerBullets = append(g.playerBullets[:i], g.playerBullets[i+1:]...)
						i--
						break
					}
				}
			}

			// 移除超出屏幕的子弹
			if i > 0 && i < len(g.playerBullets) {
				if g.playerBullets[i].x < 0 || g.playerBullets[i].x > screenWidth || g.playerBullets[i].y < 0 || g.playerBullets[i].y > screenHeight {
					g.playerBullets = append(g.playerBullets[:i], g.playerBullets[i+1:]...)
					i--
				}
			}
		}
	}

	// 更新敌人坦克状态
	for i := 0; i < len(g.enemyTanks); i++ {
		// 简单的随机移动逻辑
		if int(ebiten.ActualTPS())%changeDirInterval == 0 && !g.enemyTanks[i].directionChanged {
			g.enemyTanks[i].direction = rand.Intn(4)
			g.enemyTanks[i].directionChanged = true
		} else if int(ebiten.ActualTPS())%changeDirInterval != 0 {
			g.enemyTanks[i].directionChanged = false
		}

		switch g.enemyTanks[i].direction {
		case 0:
			if g.enemyTanks[i].y > 0 {
				g.enemyTanks[i].y -= tankSpeed
			}
		case 1:
			if g.enemyTanks[i].x < screenWidth-20 {
				g.enemyTanks[i].x += tankSpeed
			}
		case 2:
			if g.enemyTanks[i].y < screenHeight-20 {
				g.enemyTanks[i].y += tankSpeed
			}
		case 3:
			if g.enemyTanks[i].x > 0 {
				g.enemyTanks[i].x -= tankSpeed
			}
		}

		// 简单的随机射击逻辑
		if int(ebiten.ActualTPS())%shootInterval == 0 && !g.enemyTanks[i].hasShot {
			bullet := Bullet{
				x:         g.enemyTanks[i].x + 8,
				y:         g.enemyTanks[i].y + 8,
				direction: g.enemyTanks[i].direction,
			}
			g.enemyBullets = append(g.enemyBullets, bullet)
			g.enemyTanks[i].hasShot = true
		} else if int(ebiten.ActualTPS())%shootInterval != 0 {
			g.enemyTanks[i].hasShot = false
		}

		// 更新子弹位置
		for i := 0; i < len(g.enemyBullets); i++ {
			switch g.enemyBullets[i].direction {
			case 0:
				g.enemyBullets[i].y -= bulletSpeed
			case 1:
				g.enemyBullets[i].x += bulletSpeed
			case 2:
				g.enemyBullets[i].y += bulletSpeed
			case 3:
				g.enemyBullets[i].x -= bulletSpeed
			}

			// 检测敌方子弹与玩家坦克的碰撞
			if g.playerTank != nil && len(g.enemyBullets) > 0 {
				if checkCollision(g.enemyBullets[i].x, g.enemyBullets[i].y, 5, 5, g.playerTank.x, g.playerTank.y, 20, 20) {
					// 移除玩家坦克
					g.playerTank = nil
					// 移除子弹
					g.enemyBullets = append(g.enemyBullets[:i], g.enemyBullets[i+1:]...)
					i--
					continue
				}
			}

			// 检测子弹与墙的碰撞
			for j := 0; j < len(g.walls); j++ {
				if checkCollision(g.enemyBullets[i].x, g.enemyBullets[i].y, 5, 5, g.walls[j].x, g.walls[j].y, g.walls[j].width, g.walls[j].height) {
					g.walls[j].health--
					if g.walls[j].health <= 0 {
						g.walls = append(g.walls[:j], g.walls[j+1:]...)
					}
					g.enemyBullets = append(g.enemyBullets[:i], g.enemyBullets[i+1:]...)
					i--
					break
				}
			}

			// 移除超出屏幕的子弹
			if i > 0 && i < len(g.enemyBullets) {
				if g.enemyBullets[i].x < 0 || g.enemyBullets[i].x > screenWidth || g.enemyBullets[i].y < 0 || g.enemyBullets[i].y > screenHeight {
					g.enemyBullets = append(g.enemyBullets[:i], g.enemyBullets[i+1:]...)
					i--
				}
			}
		}
	}

	return nil
}

// Draw 绘制游戏画面
func (g *Game) Draw(screen *ebiten.Image) {
	// 绘制玩家坦克
	if g.playerTank != nil {
		ebitenutil.DrawRect(screen, g.playerTank.x, g.playerTank.y, 20, 20, color.RGBA{0, 255, 0, 255})
		switch g.playerTank.direction {
		case 0:
			ebitenutil.DrawLine(screen, g.playerTank.x+10, g.playerTank.y, g.playerTank.x+10, g.playerTank.y-10, color.RGBA{0, 255, 0, 255})
		case 1:
			ebitenutil.DrawLine(screen, g.playerTank.x+20, g.playerTank.y+10, g.playerTank.x+30, g.playerTank.y+10, color.RGBA{0, 255, 0, 255})
		case 2:
			ebitenutil.DrawLine(screen, g.playerTank.x+10, g.playerTank.y+20, g.playerTank.x+10, g.playerTank.y+30, color.RGBA{0, 255, 0, 255})
		case 3:
			ebitenutil.DrawLine(screen, g.playerTank.x, g.playerTank.y+10, g.playerTank.x-10, g.playerTank.y+10, color.RGBA{0, 255, 0, 255})
		}
	}

	// 绘制玩家子弹
	for _, bullet := range g.playerBullets {
		ebitenutil.DrawRect(screen, bullet.x, bullet.y, 5, 5, color.RGBA{0, 255, 0, 255})
	}

	// 绘制敌人坦克
	for _, enemyTank := range g.enemyTanks {
		ebitenutil.DrawRect(screen, enemyTank.x, enemyTank.y, 20, 20, color.RGBA{255, 0, 0, 255})
		switch enemyTank.direction {
		case 0:
			ebitenutil.DrawLine(screen, enemyTank.x+10, enemyTank.y, enemyTank.x+10, enemyTank.y-10, color.RGBA{255, 0, 0, 255})
		case 1:
			ebitenutil.DrawLine(screen, enemyTank.x+20, enemyTank.y+10, enemyTank.x+30, enemyTank.y+10, color.RGBA{255, 0, 0, 255})
		case 2:
			ebitenutil.DrawLine(screen, enemyTank.x+10, enemyTank.y+20, enemyTank.x+10, enemyTank.y+30, color.RGBA{255, 0, 0, 255})
		case 3:
			ebitenutil.DrawLine(screen, enemyTank.x, enemyTank.y+10, enemyTank.x-10, enemyTank.y+10, color.RGBA{255, 0, 0, 255})
		}
	}

	// 绘制敌人子弹
	for _, bullet := range g.enemyBullets {
		ebitenutil.DrawRect(screen, bullet.x, bullet.y, 5, 5, color.RGBA{255, 0, 0, 255})
	}

	// 绘制墙壁
	for _, wall := range g.walls {
		ebitenutil.DrawRect(screen, wall.x, wall.y, wall.width, wall.height, color.RGBA{128, 128, 128, 255})
	}
}

// Layout 返回游戏画面的布局
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
