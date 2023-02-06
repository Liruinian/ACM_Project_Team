function typer() {
  const container = document.getElementById("content");
  const data =
    "WARNING: terminal is not fully functional~[  0.000000] Initializing cgroup subsys cpuset~[  0.000000] Initializing cgroup subsys cpu~[  0.000000] Linux version 2.6.32-21-generic (buildd@rothera) (gcc version 4.4.3 (Ubuntu 4.4.3-4ubuntu5) ) #32-Ubuntu SMP Fri Apr 16 08:10:02 UTC 2010 (Ubuntu 2.6.32-21.32-generic 2.6.32.11+drm33.2)~[  0.000000] KERNEL supported cpus:~[  0.000000]  Intel GenuineIntel~[  0.000000]  AMD AuthenticAMD~[  0.000000]  NSC Geode by NSC~[  0.000000]  Cyrix CyrixInstead~[  0.000000]  Centaur CentaurHauls~[  0.000000]  Transmeta GenuineTMx86~[  0.000000]  Transmeta TransmetaCPU~[  0.000000]  UMC UMC UMC UMC~[  0.000000] BIOS-provided physical RAM map:~[  0.000000] BIOS-e820: 0000000000000000 - 000000000009f800 (usable)~[  0.000000] BIOS-e820: 000000000009f800 - 00000000000a0000 (reserved)~[  0.000000] BIOS-e820: 00000000000ca000 - 00000000000cc000 (reserved)~[  0.000000] BIOS-e820: 00000000000dc000 - 00000000000e0000 (reserved)~[  0.000000] BIOS-e820: 00000000000e4000 - 0000000000100000 (reserved)~[  0.000000] BIOS-e820: 0000000000100000 - 000000003fef0000 (usable)~[  0.000000] BIOS-e820: 000000003fef0000 - 000000003feff000 (ACPI data)~[  0.000000] BIOS-e820: 000000003feff000 - 000000003ff00000 (ACPI NVS)".split(
      ""
    );
  let index = 0;
  function writing() {
    if (index < data.length) {
      if (data[index] == "~") {
        container.innerHTML += "<br />";
      }
      container.innerHTML += data[index++];
      requestAnimationFrame(writing);
    }
  }
  writing();
}

window.onload = () => {
  typer();
};
