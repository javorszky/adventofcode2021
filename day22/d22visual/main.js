import './style.css'
import * as THREE from 'three'
import { TrackballControls } from 'three/examples/jsm/controls/TrackballControls'

// document.querySelector('#app').innerHTML = `
//   <h1>Hello Vite!</h1>
//   <a href="https://vitejs.dev/guide/features.html" target="_blank">Documentation</a>
// `

const randomColor = function () {
  var rgb = new Uint32Array(1)
  self.crypto.getRandomValues(rgb)

  console.log(rgb)

  return rgb[0]
}

const width = window.innerWidth
const height = window.innerHeight

// Create the scene
const scene = new THREE.Scene()

// The camera
const camera = new THREE.PerspectiveCamera(35, width / height, 1, 1000)

const renderer = new THREE.WebGLRenderer()
renderer.setSize(window.innerWidth, window.innerHeight)
document.querySelector('#app').appendChild(renderer.domElement)

const geometry = new THREE.BoxGeometry()
const material = new THREE.MeshBasicMaterial({ color: 0x00ff00 })
const cube = new THREE.Mesh(geometry, material)
scene.add(cube)
scene.add(new THREE.AxesHelper(55))

camera.position.z = 150
camera.position.x = 150
camera.position.y = 150
camera.lookAt(0, 0, 0)

const controls = new TrackballControls(camera, renderer.domElement)

function parseCubes (cubeString) {
  //   const cubes = JSON.parse(cubeString)
  //   console.log(cubeString)

  for (const [key, value] of Object.entries(cubeString)) {
    const x = value.x_to - value.x_from + 1
    const y = value.y_to - value.y_from + 1
    const z = value.z_to - value.z_from + 1

    const geo = new THREE.BoxGeometry(x, y, z)

    const material = new THREE.MeshBasicMaterial({
      color: randomColor(),
      opacity: 0.3,
      transparent: true
    })
    const cube = new THREE.Mesh(geo, material)

    console.log(cube.position)
    cube.position.x = value.x_from + x / 2
    cube.position.y = value.y_from + y / 2
    cube.position.z = value.z_from + z / 2

    scene.add(cube)
  }
}

function animate () {
  requestAnimationFrame(animate)
  renderer.render(scene, camera)
  controls.update()
}
animate()

const bla = {
  '10/10/10/10/10/10/on': {
    x_from: 10,
    x_to: 10,
    y_from: 10,
    y_to: 10,
    z_from: 10,
    z_to: 10,
    Flip: 1
  },
  '10/10/11/12/12/12/on': {
    x_from: 10,
    x_to: 10,
    y_from: 11,
    y_to: 12,
    z_from: 12,
    z_to: 12,
    Flip: 1
  },
  '10/10/12/12/10/11/on': {
    x_from: 10,
    x_to: 10,
    y_from: 12,
    y_to: 12,
    z_from: 10,
    z_to: 11,
    Flip: 1
  },
  '10/12/10/10/12/12/on': {
    x_from: 10,
    x_to: 12,
    y_from: 10,
    y_to: 10,
    z_from: 12,
    z_to: 12,
    Flip: 1
  },
  '11/11/11/13/12/13/on': {
    x_from: 11,
    x_to: 11,
    y_from: 11,
    y_to: 13,
    z_from: 12,
    z_to: 13,
    Flip: 1
  },
  '11/12/12/12/10/10/on': {
    x_from: 11,
    x_to: 12,
    y_from: 12,
    y_to: 12,
    z_from: 10,
    z_to: 10,
    Flip: 1
  },
  '11/13/12/13/11/11/on': {
    x_from: 11,
    x_to: 13,
    y_from: 12,
    y_to: 13,
    z_from: 11,
    z_to: 11,
    Flip: 1
  },
  '12/12/10/10/11/11/on': {
    x_from: 12,
    x_to: 12,
    y_from: 10,
    y_to: 10,
    z_from: 11,
    z_to: 11,
    Flip: 1
  },
  '12/12/10/11/10/10/on': {
    x_from: 12,
    x_to: 12,
    y_from: 10,
    y_to: 11,
    z_from: 10,
    z_to: 10,
    Flip: 1
  },
  '12/13/11/11/11/13/on': {
    x_from: 12,
    x_to: 13,
    y_from: 11,
    y_to: 11,
    z_from: 11,
    z_to: 13,
    Flip: 1
  },
  '12/13/12/13/12/13/on': {
    x_from: 12,
    x_to: 13,
    y_from: 12,
    y_to: 13,
    z_from: 12,
    z_to: 13,
    Flip: 1
  }
}

parseCubes(bla)
