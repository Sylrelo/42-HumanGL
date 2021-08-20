package main

func createAnimation(animation *Animation, start, end, bodyPart int, rot Vec3f32) {
	(*animation).keyframes = append((*animation).keyframes, AnimationDetail{
		BodyPart: bodyPart,
		rotation: rot,
		start:    start,
		end:      end,
	})
}
func createAnimationTranslation(animation *AnimationTranslation, start, end int, trans Vec3f32) {
	(*animation).keyframes = append((*animation).keyframes, AnimationDetailTranslation{
		start:       start,
		end:         end,
		translation: trans,
	})
}

func createJumpingTranslation() AnimationTranslation {
	var animation AnimationTranslation
	animation.duration = 500

	createAnimationTranslation(&animation, 0, 40, Vec3f32{0, -220, 0})
	createAnimationTranslation(&animation, 60, 80, Vec3f32{0, 500, 0})

	// createAnimationTranslation(&animation, 80, 160, Vec3f32{0, -520, 0})
	// createAnimationTranslation(&animation, 60, 120, Vec3f32{0, 500, 0})

	//createAnimationTranslation(&animation, 120, 180, Vec3f32{0, -145, 0})
	//createAnimationTranslation(&animation, 180, 220, Vec3f32{0, -145, 0})

	return animation
}

func createJumpingAnimation() Animation {
	var animation Animation

	animation.duration = 500

	// Go down \\
	createAnimation(&animation, 0, 40, TORSO, Vec3f32{-25, 0, 0})

	createAnimation(&animation, 0, 40, HEAD, Vec3f32{25, 0, 0})

	createAnimation(&animation, 0, 40, LULEG, Vec3f32{90, 0, 0})
	createAnimation(&animation, 0, 40, LLLEG, Vec3f32{-135, 0, 0})

	createAnimation(&animation, 0, 40, RULEG, Vec3f32{90, 0, 0})
	createAnimation(&animation, 0, 40, RLLEG, Vec3f32{-135, 0, 0})

	createAnimation(&animation, 0, 40, RUARM, Vec3f32{10, 0, 0})
	createAnimation(&animation, 0, 40, LUARM, Vec3f32{10, 0, 0})

	createAnimation(&animation, 0, 40, RLARM, Vec3f32{45, 0, 0})
	createAnimation(&animation, 0, 40, LLARM, Vec3f32{45, 0, 0})

	// Jump \\
	createAnimation(&animation, 60, 75, TORSO, Vec3f32{15, 0, 0})

	createAnimation(&animation, 60, 75, HEAD, Vec3f32{-15, 0, 0})

	createAnimation(&animation, 60, 75, LULEG, Vec3f32{-90, 0, 0})
	createAnimation(&animation, 60, 75, LLLEG, Vec3f32{135, 0, 0})

	createAnimation(&animation, 60, 75, RULEG, Vec3f32{-90, 0, 0})
	createAnimation(&animation, 60, 75, RLLEG, Vec3f32{135, 0, 0})

	createAnimation(&animation, 60, 75, RUARM, Vec3f32{-10, 0, 0})
	createAnimation(&animation, 60, 75, LUARM, Vec3f32{-10, 0, 0})

	createAnimation(&animation, 60, 75, RLARM, Vec3f32{-45, 0, 0})
	createAnimation(&animation, 60, 75, LLARM, Vec3f32{-45, 0, 0})

	// Go ball \\
	createAnimation(&animation, 80, 120, TORSO, Vec3f32{15, 0, 0})

	createAnimation(&animation, 80, 120, HEAD, Vec3f32{-15, 0, 0})

	createAnimation(&animation, 80, 120, LULEG, Vec3f32{90, 0, 0})
	createAnimation(&animation, 80, 120, LLLEG, Vec3f32{-135, 0, 0})

	createAnimation(&animation, 80, 120, RULEG, Vec3f32{90, 0, 0})
	createAnimation(&animation, 80, 120, RLLEG, Vec3f32{-135, 0, 0})

	createAnimation(&animation, 80, 120, RUARM, Vec3f32{10, 0, 0})
	createAnimation(&animation, 80, 120, LUARM, Vec3f32{10, 0, 0})

	createAnimation(&animation, 80, 120, RLARM, Vec3f32{45, 0, 0})
	createAnimation(&animation, 80, 120, LLARM, Vec3f32{45, 0, 0})

	// Prepare falling \\
	createAnimation(&animation, 120, 160, TORSO, Vec3f32{5, 0, 0})

	createAnimation(&animation, 120, 160, HEAD, Vec3f32{-5, 0, 0})

	createAnimation(&animation, 120, 160, LULEG, Vec3f32{-70, 0, 0})
	createAnimation(&animation, 120, 160, LLLEG, Vec3f32{120, 0, 0})

	createAnimation(&animation, 120, 160, RULEG, Vec3f32{-70, 0, 0})
	createAnimation(&animation, 120, 160, RLLEG, Vec3f32{120, 0, 0})

	createAnimation(&animation, 120, 160, RUARM, Vec3f32{-10, 0, 0})
	createAnimation(&animation, 120, 160, LUARM, Vec3f32{-10, 0, 0})

	createAnimation(&animation, 120, 160, RLARM, Vec3f32{-45, 0, 0})
	createAnimation(&animation, 120, 160, LLARM, Vec3f32{-45, 0, 0})

	// Get on ground \\
	createAnimation(&animation, 160, 180, TORSO, Vec3f32{-25, 0, 0})

	createAnimation(&animation, 160, 180, HEAD, Vec3f32{25, 0, 0})

	createAnimation(&animation, 160, 180, LULEG, Vec3f32{50, 0, 0})
	createAnimation(&animation, 160, 180, LLLEG, Vec3f32{-100, 0, 0})

	createAnimation(&animation, 160, 180, RULEG, Vec3f32{50, 0, 0})
	createAnimation(&animation, 160, 180, RLLEG, Vec3f32{-100, 0, 0})

	createAnimation(&animation, 160, 180, RUARM, Vec3f32{10, 0, 0})
	createAnimation(&animation, 160, 180, LUARM, Vec3f32{10, 0, 0})

	createAnimation(&animation, 160, 180, RLARM, Vec3f32{45, 0, 0})
	createAnimation(&animation, 160, 180, LLARM, Vec3f32{45, 0, 0})

	// Get back up \\
	createAnimation(&animation, 180, 220, TORSO, Vec3f32{15, 0, 0})

	createAnimation(&animation, 180, 220, HEAD, Vec3f32{-15, 0, 0})

	createAnimation(&animation, 180, 220, LULEG, Vec3f32{-70, 0, 0})
	createAnimation(&animation, 180, 220, LLLEG, Vec3f32{115, 0, 0})

	createAnimation(&animation, 180, 220, RULEG, Vec3f32{-70, 0, 0})
	createAnimation(&animation, 180, 220, RLLEG, Vec3f32{115, 0, 0})

	createAnimation(&animation, 180, 220, RUARM, Vec3f32{-10, 0, 0})
	createAnimation(&animation, 180, 220, LUARM, Vec3f32{-10, 0, 0})

	createAnimation(&animation, 180, 220, RLARM, Vec3f32{-45, 0, 0})
	createAnimation(&animation, 180, 220, LLARM, Vec3f32{-45, 0, 0})

	// dab \\
	createAnimation(&animation, 230, 245, LUARM, Vec3f32{75, 75, 0})
	createAnimation(&animation, 230, 245, LLARM, Vec3f32{90, 90, 65})
	createAnimation(&animation, 230, 245, RUARM, Vec3f32{105, -105, 0})
	createAnimation(&animation, 230, 245, HEAD, Vec3f32{-35, 45, 0})

	createAnimation(&animation, 300, 360, LUARM, Vec3f32{-75, -75, 0})
	createAnimation(&animation, 300, 360, LLARM, Vec3f32{-90, -90, -65})
	createAnimation(&animation, 300, 360, RUARM, Vec3f32{-105, 105, 0})
	createAnimation(&animation, 300, 360, HEAD, Vec3f32{35, -45, 0})

	//createAnimation(&animation, 60, 160, TORSO, Vec3f32{360, 0, 0})

	return animation
}

func createWalkingAnimation() Animation {
	var animation Animation

	animation.duration = 220

	// return animation
	createAnimation(&animation, 0, 0, LULEG, Vec3f32{-30 * 0.5, 0, 0}) // mov_1_left_leg
	createAnimation(&animation, 0, 0, LLLEG, Vec3f32{-15 * 0.5, 0, 0}) // mov_1_left_leg

	createAnimation(&animation, 0, 40, LULEG, Vec3f32{30 * 0.5, 0, 0})  // mov_2_left_leg
	createAnimation(&animation, 0, 40, LLLEG, Vec3f32{-45 * 0.5, 0, 0}) // mov_2_left_leg

	createAnimation(&animation, 40, 80, LULEG, Vec3f32{30 * 0.5, 0, 0})  // mov_3_left_leg
	createAnimation(&animation, 40, 80, LLLEG, Vec3f32{-15 * 0.5, 0, 0}) // mov_3_left_leg

	createAnimation(&animation, 80, 120, LULEG, Vec3f32{15 * 0.5, 0, 0}) // mov_4_left_leg
	createAnimation(&animation, 80, 120, LLLEG, Vec3f32{30 * 0.5, 0, 0}) // mov_4_left_leg

	createAnimation(&animation, 120, 160, LULEG, Vec3f32{-30 * 0.5, 0, 0}) // mov_5_left_leg
	createAnimation(&animation, 120, 160, LLLEG, Vec3f32{30 * 0.5, 0, 0})  // mov_5_left_leg

	createAnimation(&animation, 160, 180, LULEG, Vec3f32{-15 * 0.5, 0, 0}) // mov_6_left_leg
	createAnimation(&animation, 160, 180, LLLEG, Vec3f32{15 * 0.5, 0, 0})  // mov_6_left_leg

	createAnimation(&animation, 180, 220, LULEG, Vec3f32{-30 * 0.5, 0, 0}) // mov_7_left_leg
	createAnimation(&animation, 180, 220, LLLEG, Vec3f32{-15 * 0.5, 0, 0}) // mov_7_left_leg

	createAnimation(&animation, 0, 0, RULEG, Vec3f32{45 * 0.5, 0, 0})  // mov_1_right_leg
	createAnimation(&animation, 0, 0, RLLEG, Vec3f32{-45 * 0.5, 0, 0}) // mov_1_right_leg

	createAnimation(&animation, 0, 40, RULEG, Vec3f32{-30 * 0.5, 0, 0}) // mov_2_right_leg
	createAnimation(&animation, 0, 40, RLLEG, Vec3f32{30 * 0.5, 0, 0})  // mov_2_right_leg

	createAnimation(&animation, 40, 60, RULEG, Vec3f32{-15 * 0.5, 0, 0}) // mov_3_right_leg
	createAnimation(&animation, 40, 60, RLLEG, Vec3f32{15 * 0.5, 0, 0})  // mov_3_right_leg

	createAnimation(&animation, 60, 120, RULEG, Vec3f32{-30 * 0.5, 0, 0}) // mov_4_right_leg
	createAnimation(&animation, 60, 120, RLLEG, Vec3f32{-15 * 0.5, 0, 0}) // mov_4_right_leg

	createAnimation(&animation, 120, 160, RULEG, Vec3f32{30 * 0.5, 0, 0})  // mov_5_right_leg
	createAnimation(&animation, 120, 160, RLLEG, Vec3f32{-45 * 0.5, 0, 0}) // mov_5_right_leg

	createAnimation(&animation, 160, 190, RULEG, Vec3f32{30 * 0.5, 0, 0})  // mov_6_right_leg
	createAnimation(&animation, 160, 190, RLLEG, Vec3f32{-15 * 0.5, 0, 0}) // mov_6_right_leg

	createAnimation(&animation, 190, 220, RULEG, Vec3f32{15 * 0.5, 0, 0}) // mov_7_right_leg
	createAnimation(&animation, 190, 220, RLLEG, Vec3f32{30 * 0.5, 0, 0}) // mov_7_right_leg

	createAnimation(&animation, 0, 0, RUARM, Vec3f32{-45, 0, 0})     // mov_1_left_arm
	createAnimation(&animation, 0, 0, RLARM, Vec3f32{45, 0, 0})      // mov_1_left_arm
	createAnimation(&animation, 0, 120, RUARM, Vec3f32{70, 0, 0})    // mov_2_left_arm
	createAnimation(&animation, 120, 220, RUARM, Vec3f32{-70, 0, 0}) // mov_3_left_arm

	createAnimation(&animation, 0, 0, LUARM, Vec3f32{25, 0, 0})     // mov_1_right_arm
	createAnimation(&animation, 0, 0, LLARM, Vec3f32{45, 0, 0})     // mov_1_right_arm
	createAnimation(&animation, 0, 120, LUARM, Vec3f32{-70, 0, 0})  // mov_2_right_arm
	createAnimation(&animation, 120, 220, LUARM, Vec3f32{70, 0, 0}) // mov_3_right_arm

	createAnimation(&animation, 0, 0, TORSO, Vec3f32{0, 25, 0})    // mov_1_torso
	createAnimation(&animation, 0, 120, TORSO, Vec3f32{0, -50, 0}) // mov_2_torso
	createAnimation(&animation, 0, 220, TORSO, Vec3f32{0, 50, 0})  // mov_3_torso

	createAnimation(&animation, 0, 0, HEAD, Vec3f32{0, -25, 0})   // mov_1_head
	createAnimation(&animation, 0, 120, HEAD, Vec3f32{0, 50, 0})  // mov_2_head
	createAnimation(&animation, 0, 220, HEAD, Vec3f32{0, -50, 0}) // mov_3_head

	return animation
}

func createFuckUAnimation() Animation {
	var animation Animation

	animation.duration = 240

	// return animation
	createAnimation(&animation, 0, 40, LUARM, Vec3f32{75, -40, 0})
	createAnimation(&animation, 0, 40, LLARM, Vec3f32{10, 0, 50})
	createAnimation(&animation, 0, 40, RUARM, Vec3f32{45, 0, 0})
	createAnimation(&animation, 20, 60, RLARM, Vec3f32{100, 0, 0})

	createAnimation(&animation, 120, 180, LUARM, Vec3f32{-75, 40, 0})
	createAnimation(&animation, 120, 175, LLARM, Vec3f32{-10, 0, -50})
	createAnimation(&animation, 100, 170, RUARM, Vec3f32{-45, 0, 0})
	createAnimation(&animation, 100, 165, RLARM, Vec3f32{-100, 0, 0})
	return animation

}

func createDabAnimation() Animation {
	var animation Animation
	animation.duration = 240

	// return animation
	createAnimation(&animation, 0, 25, LUARM, Vec3f32{75, 75, 0})
	createAnimation(&animation, 0, 25, LLARM, Vec3f32{90, 90, 65})
	createAnimation(&animation, 0, 25, RUARM, Vec3f32{105, -105, 0})
	createAnimation(&animation, 0, 25, HEAD, Vec3f32{-35, 45, 0})

	createAnimation(&animation, 140, 200, LUARM, Vec3f32{-75, -75, 0})
	createAnimation(&animation, 140, 200, LLARM, Vec3f32{-90, -90, -65})
	createAnimation(&animation, 140, 200, RUARM, Vec3f32{-105, 105, 0})
	createAnimation(&animation, 140, 200, HEAD, Vec3f32{35, -45, 0})
	return animation

}
