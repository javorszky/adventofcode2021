import "./style.css";
import * as THREE from "three";
import { TrackballControls } from "three/examples/jsm/controls/TrackballControls";
import * as cubes from "./js/cubes.js";

// document.querySelector('#app').innerHTML = `
//   <h1>Hello Vite!</h1>
//   <a href="https://vitejs.dev/guide/features.html" target="_blank">Documentation</a>
// `
let INTERSECTED;
let width, height;
let camera, scene, raycaster, renderer, controls;

const pointer = new THREE.Vector2();

init();
console.log(cubes.default);
parseCubes(cubes.default);

animate();

function init() {
  width = window.innerWidth;
  height = window.innerHeight;

  // Create the scene
  scene = new THREE.Scene();
  scene.background = new THREE.Color(0xf0f0f0);

  // The camera
  camera = new THREE.PerspectiveCamera(35, width / height, 1, 1000);
  camera.position.z = 150;
  camera.position.x = 150;
  camera.position.y = 150;
  camera.lookAt(0, 0, 0);

  // The Raycaster
  raycaster = new THREE.Raycaster();

  // The renderer
  renderer = new THREE.WebGLRenderer();
  renderer.setPixelRatio(window.devicePixelRatio);
  renderer.setSize(window.innerWidth, window.innerHeight);

  // The Controls
  controls = new TrackballControls(camera, renderer.domElement);

  document.querySelector("#app").appendChild(renderer.domElement);

  document.addEventListener("mousemove", onPointerMove);

  scene.add(new THREE.AxesHelper(55));
}

function onPointerMove(event) {
  pointer.x = (event.clientX / window.innerWidth) * 2 - 1;
  pointer.y = -(event.clientY / window.innerHeight) * 2 + 1;
}

function parseCubes(cubeString) {
  for (const [key, value] of Object.entries(cubeString)) {
    const x = value.x_to - value.x_from + 1;
    const y = value.y_to - value.y_from + 1;
    const z = value.z_to - value.z_from + 1;

    const geo = new THREE.BoxGeometry(x, y, z);

    const material = new THREE.MeshStandardMaterial({
      color: Math.random() * 0xffffff,
      opacity: 0.3,
      transparent: true,
    });
    const cube = new THREE.Mesh(geo, material);
    cube.name = key;

    console.log(cube.position);
    cube.position.x = value.x_from + x / 2;
    cube.position.y = value.y_from + y / 2;
    cube.position.z = value.z_from + z / 2;

    scene.add(cube);
  }
}

function animate() {
  requestAnimationFrame(animate);

  render();
}

function render() {
  raycaster.setFromCamera(pointer, camera);

  const intersects = raycaster.intersectObjects(scene.children, false);

  if (intersects.length > 0) {
    if (INTERSECTED != intersects[0].object) {
      if (INTERSECTED && INTERSECTED.material.type == "MeshStandardMaterial") {
        INTERSECTED.material.emissive.setHex(INTERSECTED.currentHex);
      }

      if (intersects[0].object.material.type == "MeshStandardMaterial") {
        INTERSECTED = intersects[0].object;
        console.log("the material", INTERSECTED.material);

        INTERSECTED.currentHex = INTERSECTED.material.emissive.getHex();
        INTERSECTED.material.emissive.setHex(0xff0000);
      }
    }
  } else {
    if (INTERSECTED && INTERSECTED.material.type == "MeshStandardMaterial") {
      console.log("what is wrong...", INTERSECTED);
      INTERSECTED.material.emissive.setHex(INTERSECTED.currentHex);
    }

    INTERSECTED = null;
  }

  renderer.render(scene, camera);
  controls.update();
}
