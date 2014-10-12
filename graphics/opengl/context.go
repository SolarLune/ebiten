package opengl

// #cgo LDFLAGS: -framework OpenGL
//
// #include <stdlib.h>
// #include <OpenGL/gl.h>
import "C"
import (
	"github.com/hajimehoshi/go-ebiten/graphics"
	"github.com/hajimehoshi/go-ebiten/graphics/matrix"
)

func enableAlphaBlending() {
	C.glEnable(C.GL_TEXTURE_2D)
	C.glEnable(C.GL_BLEND)
}

func flush() {
	C.glFlush()
}

type Context struct {
	screenId     graphics.RenderTargetId
	defaultId    graphics.RenderTargetId
	currentId    graphics.RenderTargetId
	screenWidth  int
	screenHeight int
	screenScale  int
}

func NewContext(screenWidth, screenHeight, screenScale int) *Context {
	context := &Context{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		screenScale:  screenScale,
	}
	defaultRenderTarget := &RenderTarget{
		width:  screenWidth * screenScale,
		height: screenHeight * screenScale,
		flipY:  true,
	}
	context.defaultId = idsInstance.addRenderTarget(defaultRenderTarget)

	var err error
	context.screenId, err = idsInstance.createRenderTarget(
		screenWidth, screenHeight, graphics.FilterNearest)
	if err != nil {
		panic("initializing the offscreen failed: " + err.Error())
	}
	context.ResetOffscreen()
	context.Clear()

	enableAlphaBlending()

	return context
}

func (c *Context) Dispose() {
	// TODO: remove the default framebuffer?
	idsInstance.deleteRenderTarget(c.screenId)
}

func (c *Context) Update(draw func(graphics.Context)) {
	c.ResetOffscreen()
	c.Clear()

	draw(c)

	c.SetOffscreen(c.defaultId)
	c.Clear()

	scale := float64(c.screenScale)
	geo := matrix.IdentityGeometry()
	geo.Scale(scale, scale)
	graphics.DrawWhole(
		c.RenderTarget(c.screenId),
		c.screenWidth,
		c.screenHeight,
		geo,
		matrix.IdentityColor())

	flush()
}

func (c *Context) Clear() {
	c.Fill(0, 0, 0)
}

func (c *Context) Fill(r, g, b uint8) {
	idsInstance.fillRenderTarget(c.currentId, r, g, b)
}

func (c *Context) Texture(id graphics.TextureId) graphics.Drawer {
	return &textureWithContext{id, c}
}

func (c *Context) RenderTarget(id graphics.RenderTargetId) graphics.Drawer {
	return &textureWithContext{idsInstance.toTexture(id), c}
}

func (c *Context) ResetOffscreen() {
	c.currentId = c.screenId
}

func (c *Context) SetOffscreen(renderTargetId graphics.RenderTargetId) {
	c.currentId = renderTargetId
}

type textureWithContext struct {
	id      graphics.TextureId
	context *Context
}

func (t *textureWithContext) Draw(
	parts []graphics.TexturePart,
	geo matrix.Geometry,
	color matrix.Color) {
	idsInstance.drawTexture(t.context.currentId, t.id, parts, geo, color)
}
