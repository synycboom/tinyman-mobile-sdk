package com.example.tinyman

import android.os.Bundle
import android.util.Log
import com.google.android.material.snackbar.Snackbar
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import com.example.tinyman.databinding.FragmentBurnBinding

class BurnFragment : Fragment() {
    private var _binding: FragmentBurnBinding? = null

    // This property is only valid between onCreateView and
    // onDestroyView.
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        _binding = FragmentBurnBinding.inflate(inflater, container, false)
        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        binding.textviewOutput.text = "This example remove all liquidity with a given asset id and ALGO token"
        binding.buttonBurn.setOnClickListener {
            val asset1Id = binding.editTextAsset1Id.text
            if (asset1Id.isEmpty()) {
                Snackbar.make(view, "Asset 1 ID is empty", Snackbar.LENGTH_LONG)
                    .setAction("Action", null).show()
            } else {
                binding.buttonBurn.isEnabled = false
                binding.buttonBurn.isClickable = false
                binding.textviewOutput.text = "Loading..."
                SDKViewModel().burn(asset1Id.toString()) {
                    binding.buttonBurn.isEnabled = true
                    binding.buttonBurn.isClickable = true
                    Log.i("TINY_MAN_MOBILE_SDK", it)
                    binding.textviewOutput.text = it
                }
            }
        }
        binding.buttonGoBack.setOnClickListener {
            findNavController().navigate(R.id.action_MintFragment_to_ExampleFragment)
        }
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }
}