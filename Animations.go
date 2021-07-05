package main

func createWalkinAnimation() Animation {
	var animation Animation

	animation.duration = 240


	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: LULEG,
		rotation: Vec3f32{30, 0, 0},
		start: 0,
		end: 60,
	})

	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: RULEG,
		rotation: Vec3f32{-30, 0, 0},
		start: 0,
		end: 60,
	})

	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: LULEG,
		rotation: Vec3f32{-30, 0, 0},
		start: 60,
		end: 120,
	})

	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: RULEG,
		rotation: Vec3f32{30, 0, 0},
		start: 60,
		end: 120,
	})

	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: LULEG,
		rotation: Vec3f32{-30, 0, 0},
		start: 120,
		end: 180,
	})

	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: RULEG,
		rotation: Vec3f32{30, 0, 0},
		start: 120,
		end: 180,
	})


	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: LULEG,
		rotation: Vec3f32{30, 0, 0},
		start: 180,
		end: 240,
	})

	animation.keyframes = append(animation.keyframes, AnimationDetail{
		BodyPart: RULEG,
		rotation: Vec3f32{-30, 0, 0},
		start: 180,
		end: 240,
	})


	// animation.keyframes = append(animation.keyframes, AnimationDetail{
	// 	BodyPart: LLLEG,
	// 	rotation: Vec3f32{-5, 0, 0},
	// 	start: 60,
	// 	end: 180,
	// })

	return animation

}