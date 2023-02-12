components {
  id: "giveClicks"
  component: "/main/scripts/giveClicks.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/testAtlas.atlas\"\n"
  "default_animation: \"Nota1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.2
    y: 0.2
    z: 0.2
  }
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/assets/clicker/Sound/chime-sound-7143-_1_.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.5\n"
  "pan: 0.0\n"
  "speed: 1.0\n"
  "loopcount: 0\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
