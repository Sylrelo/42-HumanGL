package main

func createAnimation(animation *Animation, start, end, bodyPart int, rot Vec3f32) {
	(*animation).keyframes = append((*animation).keyframes, AnimationDetail{
		BodyPart: bodyPart,
		rotation: rot,
		start:    start,
		end:      end,
	})
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
