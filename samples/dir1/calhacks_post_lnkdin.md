I spent 48 hours at Cal Hacks 11.0 over the weekend with Max Kessler and Syrus Aslam. We built a proof of concept plugin for Ableton Live that uses machine learning to generate melodies and chord progressions for Live sessions.


- Max for Live as a plugin platform empowered us to prototype much faster than writing a VST3 or Audio Unit plugin in C++

- Ableton's Live API allowed us to extract and insert MIDI directly into the user's piano roll for minimal impact on creative flow (it's real)

- Running inference based on project-wide MIDI data enables the model to generate relevant melodies and chord progressions

- MIDI as JSON -> byte stream -> UDP sockets

- Sleeplessly fine-tuned skytnt/tv2o-medium inference model hosted on Hugging Face 🤗


We'll be polishing it over the next few months. If you're a producer, feel free to contact arnav@surve.dev and I'll put together a mailing list for our first packaged build.


https://lnkd.in/gn4HdP55

#ai #ml #genai #huggingface #calhacks #ucberkeley #sunnysfo
